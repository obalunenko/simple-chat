package chat

import (
	"crypto/md5" // #nosec
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

// User is the interface that describes an object of chatUser info
type User interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

// UniqueID returns user's uniqueID
func (u chatUser) UniqueID() string {
	return u.uniqueID
}

// authHandler handles the auth middleware to check authorisation for routes that need it
type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie || cookie.Value == "" {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// MustAuth authentication middleware
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{
		next: handler,
	}
}

// logInHandler handles the third-party login process
// format: /auth/{action}/{provider}
type logInHandler struct{}

func (lh *logInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) != 4 {
		// should be 4 as we expected to have all elements provider and action
		//format: /auth/{action}/{provider}
		http.Error(w, "Wrong url", http.StatusNotAcceptable)
		return
	}

	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		login(w, provider)

	case "callback":
		callback(w, r, provider)

	default:
		http.Error(w, fmt.Sprintf("Auth action [%s] not supported", action), http.StatusNotFound)
	}
}

func login(w http.ResponseWriter, provider string) {

	providerOAUTH, err := gomniauth.Provider(provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when trying to get provider [%s]: %v", provider, err), http.StatusBadRequest)
		return
	}
	loginURL, err := providerOAUTH.GetBeginAuthURL(nil, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s: %v", provider, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", loginURL)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func callback(w http.ResponseWriter, r *http.Request, provider string) {
	providerOAUTH, err := gomniauth.Provider(provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when trying to get provider [%s]: %v", provider, err), http.StatusBadRequest)
		return
	}

	creds, err := providerOAUTH.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s: %v", provider, err), http.StatusInternalServerError)
		return
	}

	u, err := providerOAUTH.GetUser(creds)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error to get u form %s: %v", provider, err), http.StatusInternalServerError)
	}
	chatU := chatUser{
		User: u,
	}

	/* #nosec */
	m := md5.New()
	_, err = io.WriteString(m, strings.ToLower(u.Email()))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to set up cookie: %v", err), http.StatusInternalServerError)
		return
	}

	uID := fmt.Sprintf("%x", m.Sum(nil))
	chatU.uniqueID = uID

	avatarURL, err := SharedAvatarServicesList.GetAvatarURL(chatU)
	if err != nil {
		log.Fatalf("failed to get vatar url from service")
	}

	authCookieValue := objx.New(map[string]interface{}{
		"user_id":    uID,
		"name":       u.Name(),
		"avatar_url": avatarURL,
	}).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})

	w.Header().Set("Location", "/chat")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// ThirdPartyLoginHandler creates handler for third-party services authentication
func ThirdPartyLoginHandler() http.Handler {
	return &logInHandler{}
}

type logOutHandler struct{}

// LogOutHandler handles logout process
func LogOutHandler() http.Handler {
	return &logOutHandler{}
}

func (lo *logOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	w.Header().Set("Location", "/chat")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
