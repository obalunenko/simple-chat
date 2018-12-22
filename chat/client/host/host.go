package host

import (
	"fmt"
	"github.com/oleg-balunenko/simple-chat/chat/message"
	"github.com/oleg-balunenko/simple-chat/chat/types"
	"io"
	"log"
	"net"
)

// TODO: implement web-socket instead of TCP connection

// RunHost takes an ip as an argument "-listen"
// and listens for connections on the ip in argument
func Run(ip string) {

	host := new(types.Client)

	err := host.SetAddress(ip)
	if err != nil {
		log.Fatal("RunHost(ip string): Error at SetAddress(ip): ", err)

	}
	err = host.SetName()
	if err != nil {
		log.Fatal("RunHost(ip string): Error at SetName(): ", err)

	}

	listener, listenerErr := net.Listen("tcp", host.AddressString())

	defer closeListening(listener)

	if listenerErr != nil {
		log.Fatal("RunHost(ip string): Error at net.Listen: ", listenerErr)
	}

	fmt.Println("Listening on: ", host.AddressString())

	conn, acceptErr := listener.Accept()
	defer closeConnection(conn)

	if acceptErr != nil {
		log.Fatal("RunHost(ip string): Error at listener.Accept(): ", acceptErr)
	}

	fmt.Println("New connection accepted: ", conn)

	for {

		handleHost(conn, host)

	}

}

func handleHost(conn io.ReadWriter, host *types.Client) {

	jsonData := message.Receive(conn)

	addressee := new(types.Client)
	err := addressee.ObjectFromJSON(jsonData)
	if err != nil {

		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Client): Error at ObjectFromJSON(jsonData): ", err)

	}
	addressee.MessageString()

	err = host.SetMessage()
	if err != nil {
		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Client): Error at SetMessage(): ", err)
	}
	err = message.Send(host, conn)
	if err != nil {

		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Client): Error at sendData(guest, conn): ", err)

	}

}

func closeListening(listener io.Closer) {

	fmt.Println("Closing the Host listener.....")
	err := listener.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	err := connection.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}
