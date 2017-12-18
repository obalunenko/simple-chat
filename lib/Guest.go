package lib

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	guest := new(chatTypes.Client)

	err := guest.SetAddress(ip)
	if err != nil {
		log.Fatal("RunGuest(ip string): Error at SetAddress(ip): ", err)

	}
	err = guest.SetName()
	if err != nil {
		log.Fatal("RunGuest(ip string): Error at SetName(): ", err)

	}

	conn, dialErr := net.Dial("tcp", guest.Address())

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("RunGuest(ip string): Error at net.Dial: ", dialErr)

	}

	for {

		handleGuest(conn, guest)
	}

}

func handleGuest(conn io.ReadWriter, guest *chatTypes.Client) {

	err := guest.SetMessage()
	if err != nil {
		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at SetMessage(): ", err)
	}

	err = sendData(guest, conn)
	if err != nil {

		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at sendData(guest, conn): ", err)

	}

	jsonData := receiveData(conn)

	addressee := new(chatTypes.Client)
	err = addressee.ObjectFromJSON(jsonData)
	if err != nil {

		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at ObjectFromJSON(jsonData): ", err)

	}
	addressee.Message()
}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	err := connection.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}
