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

	conn, err := net.Dial("tcp", g.Address())
	g.connection = conn
	if err != nil {
		log.Fatal("RunGuest(ip string): Error at net.Dial: ", err)
		return errors.Wrap(err, "client/guest: Run")
	}
	defer func() {
		if err := g.Close(); err != nil {
			log.Fatalf("Failed to close guest connection: %v", err)
		}
	}()

	for {
		if err := g.Handle(); err != nil {
			return errors.Wrap(err, "client/guest: Run()")
		}
	}

}

// Handle handles process of receiving and sending messages
func (g *Guest) Handle() error {

	err := g.message.SetMessage(g.Name, os.Stdin)
	if err != nil {
		return errors.Wrap(err, "client/guest: Handle")
	}

	err = g.message.Send(g.connection)
	if err != nil {

		return errors.Wrap(err, "client/guest: Handle")

	}

	if err := g.message.Receive(g.connection); err != nil {
		return errors.Wrap(err, "client/guest: Handle")
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
