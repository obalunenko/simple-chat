package lib

import (
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// TODO: implement receive and send messages in JSON format. JSON should contain message and name of client

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	guest := new(chatTypes.Client)

	guest.SetAddress(ip)
	guest.SetName()

	conn, dialErr := net.Dial("tcp", guest.Address())

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("Error: ", dialErr)

	}

	for {

		handleGuest(conn, guest)
	}

}

func handleGuest(conn net.Conn, guest *chatTypes.Client) {

	guest.SetMessage()

	sendData(guest, conn)

	jsonData := receiveData(guest, conn)

	fmt.Println("Received data in string: ", string(jsonData))

}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	connection.Close()

}
