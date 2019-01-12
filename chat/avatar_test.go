package chat

import (
	"fmt"
	"path"
	"testing"

	"github.com/pkg/errors"
	gomniauthtest "github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
)

func setMockUser(url string, err error) *gomniauthtest.TestUser {
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return(url, err)
	return testUser
}
func TestWithAvatarServices(t *testing.T) {
	type args struct {
		avatars []Avatar
	}
	tests := []struct {
		id   int
		name string
		args args
		want AvatarServicesList
	}{
		{
			id:   1,
			name: "WithAvatarServices returns empty if no parameters passed",
			args: args{},
			want: AvatarServicesList{},
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(fmt.Sprintf("%d_%s:", test.id, test.name), func(t *testing.T) {
			got := WithAvatarServices(test.args.avatars...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestAvatarServicesList_GetAvatarURL(t *testing.T) {
	type fields struct {
		services []Avatar
	}
	type args struct {
		u chatUser
	}
	type expError struct {
		want bool
		msg  string
	}
	tests := []struct {
		id     int
		name   string
		fields fields
		args   args
		want   string
		error  expError
	}{
		{
			id:     1,
			name:   "GetAvatarURL should return ErrNoAvatarServicesConfigured when no services configured",
			fields: fields{},
			args: args{
				u: chatUser{
					User:     setMockUser("", nil),
					uniqueID: "testID",
				},
			},
			want: "",
			error: expError{
				want: true,
				msg:  ErrNoAvatarServicesConfigured.Error(),
			},
		},
		{
			id:   2,
			name: "GetAvatarURL should return ErrNoAvatarURL when no url found in all services",
			fields: fields{
				services: []Avatar{
					UseAuthAvatar(),
					UseGravatarAvatar(),
					UseFileSystemAvatar("/avatar/", "testdata", ".jpg"),
				},
			},

			args: args{
				u: chatUser{
					User:     setMockUser("", nil),
					uniqueID: "",
				},
			},
			want: "",
			error: expError{
				want: true,
				msg:  ErrNoAvatarURL.Error(),
			},
		},
		{
			id:   3,
			name: "GetAvatarURL should return valid avatar url if in one service",
			fields: fields{
				services: []Avatar{
					UseAuthAvatar(),
					UseGravatarAvatar(),
					UseFileSystemAvatar("/avatar/", "testdata", ".jpg"),
				},
			},

			args: args{
				u: chatUser{
					User:     setMockUser("", nil),
					uniqueID: "exist",
				},
			},
			want: gravatarBaseURL + "exist",
			error: expError{
				want: false,
				msg:  "",
			},
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(fmt.Sprintf("%d_%s:", test.id, test.name), func(t *testing.T) {
			a := AvatarServicesList{
				services: test.fields.services,
			}
			got, err := a.GetAvatarURL(test.args.u)
			switch test.error.want {
			case false:
				assert.NoError(t, err)
			case true:
				assert.EqualError(t, err, test.error.msg)
			}

			if got != test.want {
				t.Errorf("AvatarServicesList.GetAvatarURL() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_authAvatar_GetAvatarURL(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		id      int
		name    string
		a       authAvatar
		args    args
		want    string
		wantErr bool
	}{
		{
			id:   1,
			name: "GetAvatarURL return error and  empty avatar url when avatar on oauth not exist",
			a:    authAvatar{},
			args: args{
				u: &chatUser{
					User:     setMockUser("", ErrNoAvatarURL),
					uniqueID: "",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			id:   2,
			name: "GetAvatarURL return valid avatar url when avatar on oauth exist",
			a:    authAvatar{},
			args: args{
				u: &chatUser{
					User: setMockUser("test-avatar-url", nil),
				},
			},
			want:    "test-avatar-url",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(fmt.Sprintf("%d_%s:", test.id, test.name), func(t *testing.T) {
			a := UseAuthAvatar()
			got, err := a.GetAvatarURL(test.args.u)
			switch test.wantErr {
			case true:
				assert.EqualError(t, errors.Cause(err), ErrNoAvatarURL.Error(),
					"AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
			case false:
				assert.NoError(t, err, "AuthAvatar.GetAvatarUR should return no error when value present")

			}

			assert.Equal(t, test.want, got, "AuthAvatar.GetAvatarUR should return correct URL")
		})
	}

}

func Test_gravatarAvatar_GetAvatarURL(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		id      int
		name    string
		g       gravatarAvatar
		args    args
		want    string
		wantErr bool
	}{
		{
			id:   1,
			name: "GetAvatarURL should return valid url for user with id",
			g:    gravatarAvatar{},
			args: args{
				u: chatUser{
					uniqueID: "gravatar-test",
				},
			},
			want:    gravatarBaseURL + "gravatar-test",
			wantErr: false,
		},
		{
			id:   2,
			name: "GetAvatarURL should return ErrNoURL and empty url for user with empty id",
			g:    gravatarAvatar{},
			args: args{
				u: chatUser{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(fmt.Sprintf("%d_%s:", test.id, test.name), func(t *testing.T) {
			g := UseGravatarAvatar()
			got, err := g.GetAvatarURL(test.args.u)
			switch test.wantErr {
			case true:
				assert.EqualError(t, errors.Cause(err), ErrNoAvatarURL.Error())
			case false:
				assert.NoError(t, err, "gravatar.GetAvatarURL should not return an error")

			}

			assert.Equalf(t, test.want, got,
				"gravatar.GetAvatarURL wrongly returned %s", got)
		})
	}

}

func Test_fileSystemAvatar_GetAvatarURL(t *testing.T) {
	type fields struct {
		urlPath   string
		path      string
		extension string
	}
	type args struct {
		u User
	}
	tests := []struct {
		id      int
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			id:   1,
			name: "GetAvatarURL should return valid url for existed file",
			fields: fields{
				urlPath:   "/avatars/",
				path:      "testdata",
				extension: ".jpg",
			},
			args: args{
				u: &chatUser{
					uniqueID: "exist",
				},
			},
			want:    path.Join("/avatars/", "exist.jpg"),
			wantErr: false,
		},
		{
			id:   2,
			name: "GetAvatarURL should return ErrNotExist and empty url for not existed file",
			fields: fields{
				urlPath:   "/avatars/",
				path:      "testdata",
				extension: ".jpg",
			},
			args: args{
				u: &chatUser{
					uniqueID: "notexist",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(fmt.Sprintf("%d_%s:", test.id, test.name), func(t *testing.T) {

			fs := UseFileSystemAvatar(test.fields.urlPath, test.fields.path, test.fields.extension)

			got, err := fs.GetAvatarURL(test.args.u)
			switch test.wantErr {
			case true:
				assert.EqualError(t, errors.Cause(err), ErrNoAvatarURL.Error())
			case false:
				assert.NoError(t, err, "GetAvatarURL should not return error")

			}

			assert.Equal(t, test.want, got, "GetAvatarURL wrongly returned %s", got)

		})
	}

}
