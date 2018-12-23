package guest

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/simple-chat/message"
)

const prefix = "client/guest"

// Guest implements Client interface
type Guest struct {
	IP         string
	Port       string
	Name       string
	message    *message.Message
	connection net.Conn
}

// New creates new Guest object
func New(ip string, port string, name string) *Guest {
	return &Guest{
		IP:      ip,
		Port:    port,
		Name:    name,
		message: &message.Message{},
	}
}

// Address returns address of guest
func (g Guest) Address() string {
	return strings.Join([]string{g.IP, g.Port}, ":")
}

// Run start chat session for guest
func (g Guest) Run() error {
	var funcName = "Run()"

	conn, err := net.Dial("tcp", g.Address())
	g.connection = conn
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}
	defer func() {
		if err := g.Close(); err != nil {
			log.Fatalf("Failed to close guest connection: %v", err)
		}
	}()

	for {
		if err := g.Handle(); err != nil {
			return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
		}
	}

}

// Handle handles process of receiving and sending messages
func (g *Guest) Handle() error {
	var funcName = "Handle()"
	err := g.message.SetMessage(g.Name, os.Stdin)
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	err = g.message.Send(g.connection)
	if err != nil {

		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))

	}

	if err := g.message.Receive(g.connection); err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	fmt.Println(g.message.String())

	return nil
}

// Close closes guest session
func (g *Guest) Close() error {

	fmt.Println("Closing the Guest connection.....")

	err := g.connection.Close()
	return err
}
