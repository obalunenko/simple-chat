package lib

import (
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// sendData sends byte array of Client to connection net.Conn
func sendData(c *chatTypes.Client, conn net.Conn) {

	jsonDataToSend := c.ObjectToJson()

	fmt.Println("Will be send: ", string(jsonDataToSend))

	conn.Write(jsonDataToSend)
}

// receiveData receives byte array from connection net.Conn
func receiveData(conn net.Conn) (jsonData []byte) {

	_, err := conn.Read(jsonData)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Message received", jsonData)
	return jsonData
}
