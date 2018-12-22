package client

import (
	"github.com/oleg-balunenko/simple-chat/client/guest"
	"github.com/oleg-balunenko/simple-chat/client/host"
)

// Client interface contract for chat clients
type Client interface {
	Run() error
	Address() string
	Handle() error
	Close() error
}

// New returns new Client type according to isHost parameter
func New(isHost bool, ip string, port string, name string) Client {

	if isHost {
		// go run  main.go  -listen <ip>
		return host.New(ip, port, name)

	}

	// go run main.go  <ip>

	return guest.New(ip, port, name)

}
