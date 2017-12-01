package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// RunHost takes an ip as an argument "-listen"
// and listens for connections on the ip in argument
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenerErr := net.Listen("tcp", ipAndPort)
	if listenerErr != nil {
		log.Fatal("Error: ", listenerErr)
	}

	fmt.Println("Listening on: ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	fmt.Println("New connection accepted: ", conn)

	for {

		handleHost(conn)

	}

}

func handleHost(conn net.Conn) {

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Fprint(conn, replyMessage)

}
