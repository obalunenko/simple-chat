package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// RunHost takes an ip as an argument "-listen"
// and listens for connections on the ip in argument
func RunHost(ip string) {

	host := new(Client)

	host.SetAddress(ip)
	host.SetName()

	listener, listenerErr := net.Listen("tcp", host.GetAddress())

	defer closeListening(listener)

	if listenerErr != nil {
		log.Fatal("Error: ", listenerErr)
	}

	fmt.Println("Listening on: ", host.GetAddress())

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

func handleHost(conn net.Conn, host *Client) {

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received from", message)

	host.SetMessage()

	fmt.Fprint(conn, host.Name()+": "+host.GetMessage())

}

func closeListening(listener net.Listener) {

	fmt.Println("Closing the Host listener.....")
	listener.Close()

}
