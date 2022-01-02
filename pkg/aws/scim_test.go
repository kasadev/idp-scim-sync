package aws

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/slashdevops/idp-scim-sync/mocks/aws"
	"github.com/stretchr/testify/assert"
)

func ReadJSONFIleAsString(t *testing.T, fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	assert.NoError(t, err)

	return string(bytes)
}

func TestNewSCIMService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("should return AWSSCIMProvider", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		got, err := NewSCIMService(mockHTTPCLient, "https://testing.com", "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("should return AWSSCIMProvider when httpClient is nil", func(t *testing.T) {
		got, err := NewSCIMService(nil, "https://testing.com", "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("should return error when url is bad formed", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		got, err := NewSCIMService(mockHTTPCLient, "https://%%testing.com", "MyToken")
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("should return error when the url is empty ", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		got, err := NewSCIMService(mockHTTPCLient, "", "MyToken")
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrURLEmpty)
		assert.Nil(t, got)
	})
}

func TestDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	endpoint := "https://testing.com"

	t.Run("should return error when error come from request", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		mockHTTPCLient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("test error"))

		got, err := NewSCIMService(mockHTTPCLient, endpoint, "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)

		req := httptest.NewRequest(http.MethodGet, endpoint, nil)

		resp, err := got.do(context.Background(), req)
		assert.Error(t, err)

		assert.Nil(t, resp)
	})

	t.Run("should return valid response", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		mockResp := &http.Response{
			Status:        "200 OK",
			StatusCode:    http.StatusOK,
			Proto:         "HTTP/1.1",
			Body:          io.NopCloser(strings.NewReader("Hello, test world!")),
			ContentLength: int64(len("Hello, test world!")),
		}

		mockHTTPCLient.EXPECT().Do(gomock.Any()).Return(mockResp, nil)

		got, err := NewSCIMService(mockHTTPCLient, endpoint, "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)

		req := httptest.NewRequest(http.MethodGet, endpoint, nil)

		resp, err := got.do(context.Background(), req)
		assert.NoError(t, err)

		assert.NotNil(t, resp)
		assert.Equal(t, mockResp, resp)
	})
}

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	endpoint := "https://testing.com"
	CreateUserResponseFile := "testdata/CreateUserResponse_Active.json"

	t.Run("should return a valid response with a valid request", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		jsonResp := ReadJSONFIleAsString(t, CreateUserResponseFile)

		httpResp := &http.Response{
			Status:     "201 OK",
			StatusCode: http.StatusCreated,
			Header: http.Header{
				"Date":             []string{"Tue, 31 Mar 2020 02:36:15 GMT"},
				"Content-Type":     []string{"application/json"},
				"x-amzn-RequestId": []string{"abbf9e53-9ecc-46d2-8efe-104a66ff128f"},
			},
			Proto:         "HTTP/1.1",
			Body:          io.NopCloser(strings.NewReader(jsonResp)),
			ContentLength: int64(len(jsonResp)),
		}

		mockHTTPCLient.EXPECT().Do(gomock.Any()).Return(httpResp, nil)

		got, err := NewSCIMService(mockHTTPCLient, endpoint, "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)

		usrr := &CreateUserRequest{
			ID:         "1",
			ExternalID: "1",
			UserName:   "user.1@mail.com",
			Name: Name{
				FamilyName: "1",
				GivenName:  "test",
			},
			DisplayName: "user 1",
			Emails: []Email{
				{
					Value:   "user.1@mail.com",
					Type:    "work",
					Primary: true,
				},
			},
			Active: true,
		}

		resp, err := got.CreateUser(context.Background(), usrr)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		assert.Equal(t, "1", resp.ID)
		assert.Equal(t, "1", resp.ExternalID)
		assert.Equal(t, "user.1@mail.com", resp.UserName)
		assert.Equal(t, "user", resp.Name.GivenName)
		assert.Equal(t, "1", resp.Name.FamilyName)
		assert.Equal(t, "user 1", resp.DisplayName)
		assert.Equal(t, "user.1@mail.com", resp.Emails[0].Value)
		assert.Equal(t, "work", resp.Emails[0].Type)
		assert.Equal(t, true, resp.Emails[0].Primary)
		assert.Equal(t, true, resp.Active)
	})
}

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	endpoint := "https://testing.com"
	reqURL, err := url.Parse(endpoint)
	assert.NoError(t, err)

	t.Run("should return a valid response with a valid request", func(t *testing.T) {
		mockHTTPCLient := mocks.NewMockHTTPClient(mockCtrl)

		userID := "1"
		reqURL.Path = path.Join(reqURL.Path, fmt.Sprintf("/Users/%s", userID))

		httpReq, err := http.NewRequestWithContext(context.Background(), "DELETE", reqURL.String(), nil)
		assert.NoError(t, err)

		httpReq.Header.Set("Accept", "application/json")
		httpReq.Header.Set("Authorization", "Bearer MyToken")

		httpResp := &http.Response{
			Status:     "204 OK",
			StatusCode: http.StatusNoContent,
			Header: http.Header{
				"Date":             []string{"Tue, 31 Mar 2020 02:36:15 GMT"},
				"Content-Type":     []string{"application/json"},
				"x-amzn-RequestId": []string{"abbf9e53-9ecc-46d2-8efe-104a66ff128f"},
			},
			Proto:         "HTTP/1.1",
			Body:          io.NopCloser(strings.NewReader("")),
			ContentLength: int64(len("")),
		}

		mockHTTPCLient.EXPECT().Do(httpReq).Return(httpResp, nil)

		got, err := NewSCIMService(mockHTTPCLient, endpoint, "MyToken")
		assert.NoError(t, err)
		assert.NotNil(t, got)

		err = got.DeleteUser(context.Background(), userID)
		assert.NoError(t, err)
	})
}