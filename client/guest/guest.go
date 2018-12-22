package guest

import (
	"fmt"
	"log"
	"net"
	"strings"

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
func (g Guest) Run() {

	conn, dialErr := net.Dial("tcp", g.Address())
	g.connection = conn
	if dialErr != nil {

		log.Fatal("RunGuest(ip string): Error at net.Dial: ", dialErr)

	}
	defer g.Close()

	for {

		g.Handle()
	}

}

// Handle handles process of receiving and sending messages
func (g *Guest) Handle() {

	err := g.message.SetMessage(g.Name)
	if err != nil {
		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Message): Error at SetMessage(): ", err)
	}

	err = g.message.Send(g.connection)
	if err != nil {

		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Message): Error at sendData(guest, conn): ", err)

	}

	g.message.Receive(g.connection)

	g.message.String()
}

// Close closes guest session
func (g *Guest) Close() {

	fmt.Println("Closing the Guest connection.....")

	err := g.connection.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}
