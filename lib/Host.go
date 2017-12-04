package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// TODO: implement receive and send messages in JSON format. JSON should contain message and name of client

// TODO: implement web-socket instead of TCP connection

// RunHost takes an ip as an argument "-listen"
// and listens for connections on the ip in argument
func RunHost(ip string) {

	host := new(chatTypes.Client)

	host.SetAddress(ip)
	host.SetName()

	listener, listenerErr := net.Listen("tcp", host.Address())

	defer closeListening(listener)

	if listenerErr != nil {
		log.Fatal("Error: ", listenerErr)
	}

	fmt.Println("Listening on: ", host.Address())

	conn, acceptErr := listener.Accept()
	defer closeConnection(conn)

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	fmt.Println("New connection accepted: ", conn)

	for {

		handleHost(conn, host)

	}

}

func handleHost(conn net.Conn, host *chatTypes.Client) {

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received from", message)

	host.SetMessage()

	fmt.Fprint(conn, host.Name()+": "+host.Message())

}

func closeListening(listener net.Listener) {

	fmt.Println("Closing the Host listener.....")
	listener.Close()

}
