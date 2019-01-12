package chat

import (
	"io/ioutil"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

var (
	// ErrNoAvatarURL is the error that is returned when the Avatar
	// instance is unable to provide an Avatar URL
	ErrNoAvatarURL = errors.New("unable to get an avatar URL")

	// ErrNoAvatarServicesConfigured is the error that is returned when
	// no no any avatar service configured at WithAvatarServices
	ErrNoAvatarServicesConfigured = errors.New("no avatar services configured")
)

// Avatar represents types capable of representing
// chatUser profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the speified client.
	GetAvatarURL(user User) (string, error)
}

// SharedAvatarServicesList keeps track of the last created AvatarServicesList
var SharedAvatarServicesList AvatarServicesList

// AvatarServicesList represents a simple AvatarServicesList that holds
// an array of services, and allows access to them.
type AvatarServicesList struct {
	services []Avatar
}

// WithAvatarServices generates a new AvatarServicesList which should be
// used to get avatar for user
func WithAvatarServices(avatars ...Avatar) AvatarServicesList {
	list := AvatarServicesList{
		services: avatars,
	}
	SharedAvatarServicesList = list
	return list
}

// GetAvatarURL gets the avatar URL from the first available service
func (a AvatarServicesList) GetAvatarURL(u chatUser) (string, error) {
	if len(a.services) == 0 {
		return "", ErrNoAvatarServicesConfigured
	}
	for _, avatar := range a.services {
		if avatarURL, err := avatar.GetAvatarURL(u); err == nil {
			return avatarURL, nil
		}

	}
	return "", ErrNoAvatarURL

}

type authAvatar struct{}

// UseAuthAvatar allow to use avatars from OAUTH
func UseAuthAvatar() Avatar {
	return authAvatar{}
}

// GetAvatarURL gets the avatar URL
func (authAvatar) GetAvatarURL(u User) (string, error) {
	avatarURL := u.AvatarURL()
	if avatarURL == "" {
		return "", errors.Wrap(ErrNoAvatarURL, "authAvatar: GetAvatarURL")
	}

	return avatarURL, nil
}

type gravatarAvatar struct{}

// UseGravatarAvatar allow to use avatars from gravatar.com service
func UseGravatarAvatar() Avatar {
	return gravatarAvatar{}
}

const gravatarBaseURL = "//www.gravatar.com/avatar/"

// GetAvatarURL gets the avatar URL
func (gravatarAvatar) GetAvatarURL(u User) (string, error) {

	if u.UniqueID() == "" {
		return "", ErrNoAvatarURL
	}

	uid, err := url.Parse(u.UniqueID())
	if err != nil {
		return "", errors.Wrap(err, "gravatarAvatar: GetAvatarURL")
	}
	base, err := url.Parse(gravatarBaseURL)
	if err != nil {
		return "", errors.Wrap(err, "gravatarAvatar: GetAvatarURL")
	}

	return base.ResolveReference(uid).String(), nil
}

type fileSystemAvatar struct {
	urlPath   string
	path      string
	extension string
}

// UseFileSystemAvatar allows to use uploaded avatars and serve them from local
// path - local path where files located
// extension - file extension
// urlPath - path that will be added to url
// template of returned avatarURL: urlPath/fileName
func UseFileSystemAvatar(urlPath string, path string, extension string) Avatar {
	return fileSystemAvatar{
		urlPath:   urlPath,
		path:      path,
		extension: extension,
	}
}

// GetAvatarURL gets the avatar URL
func (fs fileSystemAvatar) GetAvatarURL(u User) (string, error) {
	if u.UniqueID() == "" {
		return "", ErrNoAvatarURL
	}

	files, err := ioutil.ReadDir(fs.path)
	if err != nil {

		return "", errors.Wrap(errors.WithMessagef(err, "failed to open dir: %s", fs.path),
			"fileSystemAvatar: GetAvatarURL")
	}
	var filename string
	var match bool
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		match, err = path.Match(u.UniqueID()+"*", f.Name())
		if err != nil {
			return "", errors.Wrap(errors.WithMessagef(err, "failed to match file %s with id %s",
				f.Name(), u.UniqueID()), "fileSystemAvatar: GetAvatarURL")
		}
		if match {
			filename = f.Name()
			break
		}
	}

	if !match {
		return "", ErrNoAvatarURL
	}

	avatarURL := path.Join(fs.urlPath, filename)
	return avatarURL, nil

}
