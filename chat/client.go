package chat

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting chatUser.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *message
	// room is the room this client is chatting in.
	room *Room

	// userData holds information about chatUser
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()

		name, ok := c.userData["name"].(string)
		if !ok {
			log.Fatalf("failed to get name of client")
		}

		msg.Name = name
		avatarURL, exist := c.userData["avatar_url"]
		if exist {
			msg.AvatarURL, ok = avatarURL.(string)
			if !ok {
				log.Fatalf("failed to typecast avatarURL to string")
			}
		}

		c.room.forward <- msg
	}
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			return
		}
	}
}

// close called when client leave room - closes send channel and websocket connection
func (c *client) close() {
	close(c.send)
	if err := c.socket.Close(); err != nil {
		log.Fatalf("failed to close socket: %v", err)
	}
}
