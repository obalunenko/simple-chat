package host

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/simple-chat/message"
)

const prefix = "client/host"

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
func (h Host) Run() error {
	var funcName = "Run()"
	var err error

	h.listener, err = net.Listen("tcp", h.Address())
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	h.connection, err = h.listener.Accept()
	fmt.Println("Listening on: ", h.Address())
	if err != nil {

		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	defer func() {
		if err := h.Close(); err != nil {
			log.Fatalf("Failed to close host connection: %v", err)
		}
	}()

	fmt.Println("New connection accepted: ", h.connection)

	for {
		if err := h.Handle(); err != nil {
			return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
		}
	}

}

// Handle handles process of receiving and sending messages
func (h *Host) Handle() error {
	var funcName = "Handle()"

	if err := h.message.Receive(h.connection); err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	fmt.Println(h.message.String())

	err := h.message.SetMessage(h.Name, os.Stdin)
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}
	err = h.message.Send(h.connection)
	if err != nil {

		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))

	}
	return nil

}

// Close closes host session
func (h *Host) Close() error {

	fmt.Println("Closing the Guest connection.....")

	err := h.connection.Close()
	if err != nil {
		return err
	}
	fmt.Println("Closing the Host listener.....")
	err = h.listener.Close()

	return err

}
