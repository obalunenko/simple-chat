package main

import (
	"flag"
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
	"github.com/oleg-balunenko/simple-chat/web"
)

var ( // flags
	host   = flag.String("host", ":8080", "server host address")
	debug  = flag.Bool("debug", false, "debug mode enables tracing of events")
	noauth = flag.Bool("noauth", false, "allow to use chat without authentication")
)

func main() {

	printVersion()

	flag.Parse()
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New("387269925431080", "32338c322fa86dd884b72227e3303c21",
			"http://localhost:8080/auth/callback/facebook"),

		github.New("6c407c7d494c9ce3af62", "4756246093276b8320c2e7d1863b787ee915e6e2",
			"http://localhost:8080/auth/callback/github"),

		google.New("900662569273-219r26ccu7ek7tiqu26tkivo0u8dr4t9.apps.googleusercontent.com",
			"89eIA-bFrJqLbbt84ucsf86_", "http://localhost:8080/auth/callback/google"),
	)

	room := registerHandlers()
	go room.Run()

	// start the web server
	log.Printf("Starting web server on %s", *host)

	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func registerHandlers() *chat.Room {
	var room *chat.Room
	staticPath := strings.Join([]string{".", "web", "static"}, string(filepath.Separator))
	templPath := strings.Join([]string{".", "web", "templates"}, string(filepath.Separator))
	avatarsPath := strings.Join([]string{".", "web", "images", "avatars"}, string(filepath.Separator))

	// register avatar services for our chat
	chat.WithAvatarServices(
		chat.UseFileSystemAvatar("/avatars/", avatarsPath, ".jpg"),
		chat.UseAuthAvatar(),
		chat.UseGravatarAvatar())

	if *debug {
		fmt.Println("Debug mode")
		room = chat.NewRoomDebug()
	} else {
		room = chat.NewRoom()
	}

	staticHandler := web.NewFilesHandler(staticPath)
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))

	chatTemplHandler := web.NewTemplateHandler(templPath, "chat.html")

	if *noauth {
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

// google
// clientID 900662569273-219r26ccu7ek7tiqu26tkivo0u8dr4t9.apps.googleusercontent.com
// client secret 89eIA-bFrJqLbbt84ucsf86_

// facebook
// clientID 387269925431080
// clientSecret 32338c322fa86dd884b72227e3303c21

// github
// clientID 6c407c7d494c9ce3af62
// clientSecret 4756246093276b8320c2e7d1863b787ee915e6e2
