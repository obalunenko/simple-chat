package guest

import (
	"fmt"
	"github.com/oleg-balunenko/simple-chat/chat/message"
	"github.com/oleg-balunenko/simple-chat/chat/types"
	"io"
	"log"
	"net"
)

// RunGuest takes an argument ip and connects to host with ip
func Run(ip string) {

	guest := new(types.Client)

	err := guest.SetAddress(ip)
	if err != nil {
		log.Fatal("RunGuest(ip string): Error at SetAddress(ip): ", err)

	}
	err = guest.SetName()
	if err != nil {
		log.Fatal("RunGuest(ip string): Error at SetName(): ", err)

	}

	conn, dialErr := net.Dial("tcp", guest.AddressString())

	defer closeConnection(conn)

	if dialErr != nil {

		log.Fatal("RunGuest(ip string): Error at net.Dial: ", dialErr)

	}

	for {

		handleGuest(conn, guest)
	}

}

func handleGuest(conn io.ReadWriter, guest *types.Client) {

	err := guest.SetMessage()
	if err != nil {
		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at SetMessage(): ", err)
	}

	err = message.Send(guest, conn)
	if err != nil {

		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at sendData(guest, conn): ", err)

	}

	jsonData := message.Receive(conn)

	addressee := new(types.Client)
	err = addressee.ObjectFromJSON(jsonData)
	if err != nil {

		log.Fatal("handleGuest(conn net.Conn, guest *chatTypes.Client): Error at ObjectFromJSON(jsonData): ", err)

	}
	addressee.MessageString()
}

func closeConnection(connection net.Conn) {

	fmt.Println("Closing the Guest connection.....")

	err := connection.Close()
	if err != nil {
		log.Fatal("closeConnection(connection net.Conn): Error at connection.Close(): ", err)
	}

}
