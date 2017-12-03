package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/types"
)

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	guest := new(types.Client)

	guest.SetAddress(ip)
	guest.SetName()

	conn, dialErr := net.Dial("tcp", guest.GetAddress())

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("Error: ", dialErr)

	}

	for {

		handleGuest(conn, guest)
	}

}

func handleGuest(conn net.Conn, guest *types.Client) {

	guest.SetMessage()

	fmt.Fprint(conn, guest.Name()+": "+guest.GetMessage())

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Println("Message received from", replyMessage)

}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	connection.Close()

}
