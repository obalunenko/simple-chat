package chat

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"

	"github.com/oleg-balunenko/simple-chat/tracer"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// Room represents chat room implementation
type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan *message
	// join is a channel for clients wishing to join the room.
	join chan *client
	// leave is a channel for clients wishing to leave the room.
	leave chan *client
	// clients holds all current clients in this room.
	clients map[*client]bool
	// tracer will receive trace information of activity
	// in the room
	tracer tracer.Tracer
}

// NewRoom returns new Room
func NewRoom() *Room {
	return &Room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  tracer.Off(),
	}
}

// NewRoomDebug returns new Room object with enabled tracer
func NewRoomDebug() *Room {
	return &Room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  tracer.New(os.Stdout),
	}
}

// Run starts chat room
func (r *Room) Run() {
	for {
		select {
		case c := <-r.join:
			// joining to room
			r.clients[c] = true
			r.tracer.Trace("New client joined: ", c.userData["name"])
		case c := <-r.leave:
			// leaving the room
			delete(r.clients, c)
			c.close()
			r.tracer.Trace("Client left: ", c.userData["name"])
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", msg.Message, " From: ", msg.Name)
			for cl := range r.clients {
				cl.send <- msg
				r.tracer.Trace(" -- sent to client: ", cl.userData["name"])
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

// ServeHTTP implements http.Handler interface
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatalf("failed to get auth cookie: %v", err)
	}

	cl := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- cl
	defer func() {
		r.leave <- cl
	}()

	go cl.write()
	cl.read()
}
