// Code generated by MockGen. DO NOT EDIT.
// Source: internal/google/directory.go

// Package google is a generated GoMock package.
package google

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "google.golang.org/api/admin/directory/v1"
)

// MockDirectoryService is a mock of DirectoryService interface.
type MockDirectoryService struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryServiceMockRecorder
}

// MockDirectoryServiceMockRecorder is the mock recorder for MockDirectoryService.
type MockDirectoryServiceMockRecorder struct {
	mock *MockDirectoryService
}

// NewMockDirectoryService creates a new mock instance.
func NewMockDirectoryService(ctrl *gomock.Controller) *MockDirectoryService {
	mock := &MockDirectoryService{ctrl: ctrl}
	mock.recorder = &MockDirectoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectoryService) EXPECT() *MockDirectoryServiceMockRecorder {
	return m.recorder
}

// GetGroup mocks base method.
func (m *MockDirectoryService) GetGroup(groupID string) (*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", groupID)
	ret0, _ := ret[0].(*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockDirectoryServiceMockRecorder) GetGroup(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockDirectoryService)(nil).GetGroup), groupID)
}

// GetUser mocks base method.
func (m *MockDirectoryService) GetUser(userID string) (*admin.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID)
	ret0, _ := ret[0].(*admin.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDirectoryServiceMockRecorder) GetUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDirectoryService)(nil).GetUser), userID)
}

// ListGroupMembers mocks base method.
func (m *MockDirectoryService) ListGroupMembers(groupID string) ([]*admin.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupMembers", groupID)
	ret0, _ := ret[0].([]*admin.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupMembers indicates an expected call of ListGroupMembers.
func (mr *MockDirectoryServiceMockRecorder) ListGroupMembers(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMembers", reflect.TypeOf((*MockDirectoryService)(nil).ListGroupMembers), groupID)
}

// ListGroups mocks base method.
func (m *MockDirectoryService) ListGroups(query []string) ([]*admin.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", query)
	ret0, _ := ret[0].([]*admin.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockDirectoryServiceMockRecorder) ListGroups(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockDirectoryService)(nil).ListGroups), query)
}

// ListUsers mocks base method.
func (m *MockDirectoryService) ListUsers(query []string) ([]*admin.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", query)
	ret0, _ := ret[0].([]*admin.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockDirectoryServiceMockRecorder) ListUsers(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockDirectoryService)(nil).ListUsers), query)
}
