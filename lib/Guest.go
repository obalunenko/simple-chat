package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	guest := new(Client)

	guest.SetIP(ip)
	guest.SetName()

	ipAndPort := guest.IP + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("Error: ", dialErr)

	}

	for {

		handleGuest(conn, guest.Name)
	}

}

func handleGuest(conn net.Conn, guestName string) {

	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint(conn, guestName+": "+message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Println("Message received: ", replyMessage)

}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	connection.Close()

}
