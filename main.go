package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}

}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenerErr := net.Listen("tcp", ipAndPort)
	if listenerErr != nil {
		log.Fatal("Error: ", listenerErr)
	}

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: ", message)
}

func runGuest(ip string) {

}
