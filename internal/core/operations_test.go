package core

import (
	"reflect"
	"testing"

	"github.com/slashdevops/idp-scim-sync/internal/model"
	"github.com/slashdevops/idp-scim-sync/internal/utils"
)

func Test_groupsDifferences(t *testing.T) {
	type args struct {
		idp   *model.GroupsResult
		state *model.GroupsResult
	}
	tests := []struct {
		name       string
		args       args
		wantCreate *model.GroupsResult
		wantUpdate *model.GroupsResult
		wantEqual  *model.GroupsResult
		wantDelete *model.GroupsResult
	}{
		{
			name: "empty",
			args: args{
				idp: &model.GroupsResult{
					Items:     0,
					Resources: []*model.Group{},
				},
				state: &model.GroupsResult{
					Items:     0,
					Resources: []*model.Group{},
				},
			},
			wantCreate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantUpdate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantEqual: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantDelete: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
		},
		{
			name: "2 equals",
			args: args{
				idp: &model.GroupsResult{
					Items: 2,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "2@mail.com"},
					},
				},
				state: &model.GroupsResult{
					Items: 2,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "2@mail.com"},
					},
				},
			},
			wantCreate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantUpdate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantEqual: &model.GroupsResult{
				Items: 2,
				Resources: []*model.Group{
					{ID: "1", Name: "name1", Email: "1@mail.com"},
					{ID: "2", Name: "name2", Email: "2@mail.com"},
				},
			},
			wantDelete: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
		},
		{
			name: "1 equals, 1 update",
			args: args{
				idp: &model.GroupsResult{
					Items: 2,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "22@mail.com"},
					},
				},
				state: &model.GroupsResult{
					Items: 2,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "2@mail.com"},
					},
				},
			},
			wantCreate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantUpdate: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "2", Name: "name2", Email: "22@mail.com"},
				},
			},
			wantEqual: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "1", Name: "name1", Email: "1@mail.com"},
				},
			},
			wantDelete: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
		},
		{
			name: "1 equals, 1 update, 1 delete",
			args: args{
				idp: &model.GroupsResult{
					Items: 2,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "22@mail.com"},
					},
				},
				state: &model.GroupsResult{
					Items: 3,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "2@mail.com"},
						{ID: "3", Name: "name3", Email: "3@mail.com"},
					},
				},
			},
			wantCreate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantUpdate: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "2", Name: "name2", Email: "22@mail.com"},
				},
			},
			wantEqual: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "1", Name: "name1", Email: "1@mail.com"},
				},
			},
			wantDelete: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "3", Name: "name3", Email: "3@mail.com"},
				},
			},
		},
		{
			name: "1 equals, 1 update, 1 delete, 1 create",
			args: args{
				idp: &model.GroupsResult{
					Items: 4,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "22@mail.com"},
						{ID: "4", Name: "name4", Email: "4@mail.com"},
					},
				},
				state: &model.GroupsResult{
					Items: 3,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
						{ID: "2", Name: "name2", Email: "2@mail.com"},
						{ID: "3", Name: "name3", Email: "3@mail.com"},
					},
				},
			},
			wantCreate: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "4", Name: "name4", Email: "4@mail.com"},
				},
			},
			wantUpdate: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "2", Name: "name2", Email: "22@mail.com"},
				},
			},
			wantEqual: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "1", Name: "name1", Email: "1@mail.com"},
				},
			},
			wantDelete: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "3", Name: "name3", Email: "3@mail.com"},
				},
			},
		},
		{
			name: "1 update, change the ID",
			args: args{
				idp: &model.GroupsResult{
					Items: 1,
					Resources: []*model.Group{
						{ID: "11", Name: "name1", Email: "1@mail.com"},
					},
				},
				state: &model.GroupsResult{
					Items: 1,
					Resources: []*model.Group{
						{ID: "1", Name: "name1", Email: "1@mail.com"},
					},
				},
			},
			wantCreate: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantUpdate: &model.GroupsResult{
				Items: 1,
				Resources: []*model.Group{
					{ID: "11", Name: "name1", Email: "1@mail.com"},
				},
			},
			wantEqual: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
			wantDelete: &model.GroupsResult{
				Items:     0,
				Resources: []*model.Group{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCreate, gotUpdate, gotEqual, gotDelete := groupsDifferences(tt.args.idp, tt.args.state)
			if !reflect.DeepEqual(gotCreate, tt.wantCreate) {
				t.Errorf("groupsDifferences() gotCreate = %s, want %s", utils.ToJSON(gotCreate), utils.ToJSON(tt.wantCreate))
			}
			if !reflect.DeepEqual(gotUpdate, tt.wantUpdate) {
				t.Errorf("groupsDifferences() gotUpdate = %s, want %s", utils.ToJSON(gotUpdate), utils.ToJSON(tt.wantUpdate))
			}
			if !reflect.DeepEqual(gotEqual, tt.wantEqual) {
				t.Errorf("groupsDifferences() gotEqual = %s, want %s", utils.ToJSON(gotEqual), utils.ToJSON(tt.wantEqual))
			}
			if !reflect.DeepEqual(gotDelete, tt.wantDelete) {
				t.Errorf("groupsDifferences() gotDelete = %s, want %s", utils.ToJSON(gotDelete), utils.ToJSON(tt.wantDelete))
			}
		})
	}
}

func Test_usersDifferences(t *testing.T) {
	type args struct {
		idp   *model.UsersResult
		state *model.UsersResult
	}
	tests := []struct {
		name       string
		args       args
		wantCreate *model.UsersResult
		wantUpdate *model.UsersResult
		wantEqual  *model.UsersResult
		wantDelete *model.UsersResult
	}{
		{
			name: "empty",
			args: args{
				idp: &model.UsersResult{
					Items:     0,
					Resources: []*model.User{},
				},
				state: &model.UsersResult{
					Items:     0,
					Resources: []*model.User{},
				},
			},
			wantCreate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantUpdate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantEqual: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantDelete: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
		},
		{
			name: "2 equals",
			args: args{
				idp: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
					},
				},
				state: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
					},
				},
			},
			wantCreate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantUpdate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantEqual: &model.UsersResult{
				Items: 2,
				Resources: []*model.User{
					{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
					{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
				},
			},
			wantDelete: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
		},
		{
			name: "1 equals, 1 update, 1 delete",
			args: args{
				idp: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
					},
				},
				state: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
						{ID: "3", Email: "donato.ricupero@email.com", Name: model.Name{FamilyName: "donato", GivenName: "ricupero"}, DisplayName: "donato ricupero", Active: true},
					},
				},
			},
			wantCreate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantUpdate: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
				},
			},
			wantEqual: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
				},
			},
			wantDelete: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "3", Email: "donato.ricupero@email.com", Name: model.Name{FamilyName: "donato", GivenName: "ricupero"}, DisplayName: "donato ricupero", Active: true},
				},
			},
		},
		{
			name: "1 equals, 1 update",
			args: args{
				idp: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
					},
				},
				state: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
					},
				},
			},
			wantCreate: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
			wantUpdate: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
				},
			},
			wantEqual: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
				},
			},
			wantDelete: &model.UsersResult{
				Items:     0,
				Resources: []*model.User{},
			},
		},
		{
			name: "1 equals, 1 update, 1 delete, 1 create",
			args: args{
				idp: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
						{ID: "4", Email: "don.nadie@email.com", Name: model.Name{FamilyName: "don", GivenName: "nadie"}, DisplayName: "don nadie", Active: true},
					},
				},
				state: &model.UsersResult{
					Items: 2,
					Resources: []*model.User{
						{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
						{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foo", GivenName: "bar"}, DisplayName: "foo bar", Active: true},
						{ID: "3", Email: "donato.ricupero@email.com", Name: model.Name{FamilyName: "donato", GivenName: "ricupero"}, DisplayName: "donato ricupero", Active: true},
					},
				},
			},
			wantCreate: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "4", Email: "don.nadie@email.com", Name: model.Name{FamilyName: "don", GivenName: "nadie"}, DisplayName: "don nadie", Active: true},
				},
			},
			wantUpdate: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "2", Email: "foo.bar@email.com", Name: model.Name{FamilyName: "foonew", GivenName: "bar"}, DisplayName: "foonew bar", Active: true},
				},
			},
			wantEqual: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "1", Email: "john.doe@email.com", Name: model.Name{FamilyName: "john", GivenName: "doe"}, DisplayName: "john doe", Active: true},
				},
			},
			wantDelete: &model.UsersResult{
				Items: 1,
				Resources: []*model.User{
					{ID: "3", Email: "donato.ricupero@email.com", Name: model.Name{FamilyName: "donato", GivenName: "ricupero"}, DisplayName: "donato ricupero", Active: true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCreate, gotUpdate, gotEqual, gotDelete := usersDifferences(tt.args.idp, tt.args.state)
			if !reflect.DeepEqual(gotCreate, tt.wantCreate) {
				t.Errorf("usersDifferences() gotCreate = %s, want %s", utils.ToJSON(gotCreate), utils.ToJSON(tt.wantCreate))
			}
			if !reflect.DeepEqual(gotUpdate, tt.wantUpdate) {
				t.Errorf("usersDifferences() gotUpdate = %s, want %s", utils.ToJSON(gotUpdate), utils.ToJSON(tt.wantUpdate))
			}
			if !reflect.DeepEqual(gotEqual, tt.wantEqual) {
				t.Errorf("usersDifferences() gotEqual = %s, want %s", utils.ToJSON(gotEqual), utils.ToJSON(tt.wantEqual))
			}
			if !reflect.DeepEqual(gotDelete, tt.wantDelete) {
				t.Errorf("usersDifferences() gotDelete = %s, want %s", utils.ToJSON(gotDelete), utils.ToJSON(tt.wantDelete))
			}
		})
	}
}

func Test_groupsUsersDifferences(t *testing.T) {
	type args struct {
		idp   *model.GroupsUsersResult
		state *model.GroupsUsersResult
	}
	tests := []struct {
		name       string
		args       args
		wantCreate *model.GroupsUsersResult
		wantEqual  *model.GroupsUsersResult
		wantDelete *model.GroupsUsersResult
	}{
		{
			name: "empty",
			args: args{
				idp: &model.GroupsUsersResult{
					Items:     0,
					Resources: []*model.GroupUsers{},
				},
				state: &model.GroupsUsersResult{
					Items:     0,
					Resources: []*model.GroupUsers{},
				},
			},
			wantCreate: &model.GroupsUsersResult{
				Items:     0,
				Resources: []*model.GroupUsers{},
			},
			wantEqual: &model.GroupsUsersResult{
				Items:     0,
				Resources: []*model.GroupUsers{},
			},
			wantDelete: &model.GroupsUsersResult{
				Items:     0,
				Resources: []*model.GroupUsers{},
			},
		},
		{
			name: "2 equals",
			args: args{
				idp: &model.GroupsUsersResult{
					Items: 2,
					Resources: []*model.GroupUsers{
						{
							Items: 2,
							Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
								{ID: "2", Name: model.Name{FamilyName: "user", GivenName: "2"}, Email: "u.2@mail.com"},
							},
						},
						{
							Items: 1,
							Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
							},
						},
					},
				},
				state: &model.GroupsUsersResult{
					Items: 2,
					Resources: []*model.GroupUsers{
						{
							Items: 2,
							Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
								{ID: "2", Name: model.Name{FamilyName: "user", GivenName: "2"}, Email: "u.2@mail.com"},
							},
						},
						{
							Items: 1,
							Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
							},
						},
					},
				},
			},
			wantCreate: &model.GroupsUsersResult{
				Items:     0,
				Resources: []*model.GroupUsers{},
			},
			wantEqual: &model.GroupsUsersResult{
				Items: 2,
				Resources: []*model.GroupUsers{
					{
						Items: 2,
						Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
						Resources: []*model.User{
							{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
							{ID: "2", Name: model.Name{FamilyName: "user", GivenName: "2"}, Email: "u.2@mail.com"},
						},
					},
					{
						Items: 1,
						Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
						Resources: []*model.User{
							{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
						},
					},
				},
			},
			wantDelete: &model.GroupsUsersResult{
				Items:     0,
				Resources: []*model.GroupUsers{},
			},
		},
		{
			name: "1 equals, 1 create, 1 delete",
			args: args{
				idp: &model.GroupsUsersResult{
					Items: 2,
					Resources: []*model.GroupUsers{
						{
							Items: 2,
							Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
								{ID: "2", Name: model.Name{FamilyName: "user", GivenName: "2"}, Email: "u.2@mail.com"},
							},
						},
						{
							Items: 1,
							Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
							},
						},
					},
				},
				state: &model.GroupsUsersResult{
					Items: 2,
					Resources: []*model.GroupUsers{
						{
							Items: 1,
							Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
							},
						},
						{
							Items: 2,
							Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
							Resources: []*model.User{
								{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
								{ID: "3", Name: model.Name{FamilyName: "user", GivenName: "3"}, Email: "u.3@mail.com"},
							},
						},
					},
				},
			},
			wantCreate: &model.GroupsUsersResult{
				Items: 1,
				Resources: []*model.GroupUsers{
					{
						Items: 1,
						Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
						Resources: []*model.User{
							{ID: "2", Name: model.Name{FamilyName: "user", GivenName: "2"}, Email: "u.2@mail.com"},
						},
					},
				},
			},
			wantEqual: &model.GroupsUsersResult{
				Items: 2,
				Resources: []*model.GroupUsers{
					{
						Items: 1,
						Group: model.Group{ID: "1", Name: "group 1", Email: "g.1@mail.com"},
						Resources: []*model.User{
							{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
						},
					},
					{
						Items: 1,
						Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
						Resources: []*model.User{
							{ID: "1", Name: model.Name{FamilyName: "user", GivenName: "1"}, Email: "u.1@mail.com"},
						},
					},
				},
			},
			wantDelete: &model.GroupsUsersResult{
				Items: 1,
				Resources: []*model.GroupUsers{
					{
						Items: 1,
						Group: model.Group{ID: "2", Name: "group 2", Email: "g.2@mail.com"},
						Resources: []*model.User{
							{ID: "3", Name: model.Name{FamilyName: "user", GivenName: "3"}, Email: "u.3@mail.com"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCreate, gotEqual, gotDelete := groupsUsersDifferences(tt.args.idp, tt.args.state)
			if !reflect.DeepEqual(gotCreate, tt.wantCreate) {
				t.Errorf("groupsUsersDifferences() gotCreate = %s, want %s", utils.ToJSON(gotCreate), utils.ToJSON(tt.wantCreate))
			}
			if !reflect.DeepEqual(gotEqual, tt.wantEqual) {
				t.Errorf("groupsUsersDifferences() gotEqual = %s, want %s", utils.ToJSON(gotEqual), utils.ToJSON(tt.wantEqual))
			}
			if !reflect.DeepEqual(gotDelete, tt.wantDelete) {
				t.Errorf("groupsUsersDifferences() gotDelete = %s, want %s", utils.ToJSON(gotDelete), utils.ToJSON(tt.wantDelete))
			}
		})
	}
}
