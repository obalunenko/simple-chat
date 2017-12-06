package lib

import (
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// TODO: implement web-sockets instead of TCP connection

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	guest := new(chatTypes.Client)

	guest.SetAddress(ip)
	guest.SetName()

	conn, dialErr := net.Dial("tcp", guest.Address())

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("RunGuest(ip string): Error at net.Dial: ", dialErr)

	}

	for {

		handleGuest(conn, guest)
	}

}

func handleGuest(conn net.Conn, guest *chatTypes.Client) {

	guest.SetMessage()

	sendData(guest, conn)

	jsonData := receiveData(conn)

	addressee := new(chatTypes.Client)
	addressee.ObjectFromJson(jsonData)
	addressee.Message()
}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	connection.Close()

}
