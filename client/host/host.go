package host

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/oleg-balunenko/simple-chat/message"
)

// Host implements Client interface
type Host struct {
	IP         string
	Port       string
	Name       string
	message    *message.Message
	connection net.Conn
	listener   net.Listener
}

// New creates new Host object
func New(ip string, port string, name string) *Host {

	return &Host{
		IP:      ip,
		Port:    port,
		Name:    name,
		message: &message.Message{},
	}
}

// Address returns address of host
func (h Host) Address() string {
	return strings.Join([]string{h.IP, h.Port}, ":")
}

// TODO: implement web-socket instead of TCP connection

// Run start chat session for host
func (h Host) Run() {
	var err error

	h.listener, err = net.Listen("tcp", h.Address())
	if err != nil {
		log.Fatal("RunHost(ip string): Error at net.Listen: ", err)
	}

	h.connection, err = h.listener.Accept()
	fmt.Println("Listening on: ", h.Address())
	if err != nil {
		log.Fatal("RunHost(ip string): Error at listener.Accept(): ", err)
	}

	defer h.Close()
	fmt.Println("New connection accepted: ", h.connection)

	for {
		h.Handle()
	}

}

// Handle handles process of receiving and sending messages
func (h *Host) Handle() {

	h.message.Receive(h.connection)

	h.message.String()

	err := h.message.SetMessage(h.Name)
	if err != nil {
		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Message): Error at SetMessage(): ", err)
	}
	err = h.message.Send(h.connection)
	if err != nil {

		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Message): Error at sendData(guest, conn): ", err)

	}

}

// Close closes host session
func (h *Host) Close() {

	fmt.Println("Closing the Guest connection.....")

	err := h.connection.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}
	fmt.Println("Closing the Host listener.....")
	err = h.listener.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}
