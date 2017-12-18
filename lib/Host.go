package lib

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// TODO: implement web-socket instead of TCP connection

// RunHost takes an ip as an argument "-listen"
// and listens for connections on the ip in argument
func RunHost(ip string) {

	host := new(chatTypes.Client)

	err := host.SetAddress(ip)
	if err != nil {
		log.Fatal("RunHost(ip string): Error at SetAddress(ip): ", err)

	}
	err = host.SetName()
	if err != nil {
		log.Fatal("RunHost(ip string): Error at SetName(): ", err)

	}

	listener, listenerErr := net.Listen("tcp", host.Address())

	defer closeListening(listener)

	if listenerErr != nil {
		log.Fatal("RunHost(ip string): Error at net.Listen: ", listenerErr)
	}

	fmt.Println("Listening on: ", host.Address())

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

func handleHost(conn io.ReadWriter, host *chatTypes.Client) {

	jsonData := receiveData(conn)

	addressee := new(chatTypes.Client)
	err := addressee.ObjectFromJSON(jsonData)
	if err != nil {

		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Client): Error at ObjectFromJSON(jsonData): ", err)

	}
	addressee.Message()

	err = host.SetMessage()
	if err != nil {
		log.Fatal("handleHost(conn net.Conn, guest *chatTypes.Client): Error at SetMessage(): ", err)
	}
	err = sendData(host, conn)
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
