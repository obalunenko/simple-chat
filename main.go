package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"

	"github.com/oleg-balunenko/simple-chat/chat"
	"github.com/oleg-balunenko/simple-chat/config"
	"github.com/oleg-balunenko/simple-chat/web"
)

func main() {

	printVersion()

	cfg := config.Load("config.toml")

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(cfg.FacebookClientID, cfg.FacebookClientSecret,
			"http://localhost:8080/auth/callback/facebook"),

		github.New(cfg.GithubClientID, cfg.GithubClientSecret,
			"http://localhost:8080/auth/callback/github"),

		google.New(cfg.GoogleClientID,
			cfg.GoogleClientSecret, "http://localhost:8080/auth/callback/google"),
	)

	room := registerHandlers(cfg)
	go room.Run()

	// start the web server
	log.Printf("Starting web server on %s", cfg.Host)

	if err := http.ListenAndServe(cfg.Host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func registerHandlers(cfg *config.ChatConfig) *chat.Room {
	var room *chat.Room
	staticPath := strings.Join([]string{".", "web", "static"}, string(filepath.Separator))
	templPath := strings.Join([]string{".", "web", "templates"}, string(filepath.Separator))
	avatarsPath := strings.Join([]string{".", "web", "images", "avatars"}, string(filepath.Separator))

	// register avatar services for our chat
	chat.WithAvatarServices(
		chat.UseFileSystemAvatar("/avatars/", avatarsPath, ".jpg"),
		chat.UseAuthAvatar(),
		chat.UseGravatarAvatar())

	if cfg.Debug {
		fmt.Println("Debug mode")
		room = chat.NewRoomDebug()
	} else {
		room = chat.NewRoom()
	}

	staticHandler := web.NewFilesHandler(staticPath)
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))

	chatTemplHandler := web.NewTemplateHandler(templPath, "chat.html")

	if cfg.Noauth {
		http.Handle("/chat", chatTemplHandler)
	} else {
		http.Handle("/chat", chat.MustAuth(chatTemplHandler))

		loginTemplHandler := web.NewTemplateHandler(templPath, "login.html")
		http.Handle("/login", loginTemplHandler)

		http.Handle("/auth/", chat.ThirdPartyLoginHandler())
		http.Handle("/logout", chat.LogOutHandler())

		http.Handle("/upload",
			chat.MustAuth(web.NewTemplateHandler(templPath, "upload.html")))

		http.Handle("/uploader", chat.MustAuth(chat.UploaderHandler(avatarsPath)))

		avatarsHandler := web.NewFilesHandler(avatarsPath)
		http.Handle("/avatars/", chat.MustAuth(
			http.StripPrefix("/avatars/", avatarsHandler)))
	}

	http.Handle("/room", room)
	return room
}
