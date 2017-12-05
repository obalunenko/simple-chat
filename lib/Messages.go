package lib

import (
	"fmt"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// sendData sends byte array of Client to connection net.Conn
func sendData(c *chatTypes.Client, conn net.Conn) {

	jsonDataToSend := make([]byte, 500)
	jsonDataToSend = c.ObjectToJson()

	//fmt.Println("Length of []byte object:", len(jsonDataToSend))

	//fmt.Println("Will be send: ", string(jsonDataToSend))

	conn.Write(jsonDataToSend)
	//fmt.Println("Data sent...")
}

// receiveData receives byte array from connection net.Conn
func receiveData(conn net.Conn) []byte {

	dataReceived := false

	jsonDataTemp := make([]byte, 1000)

	for !dataReceived {
		_, err := conn.Read(jsonDataTemp)
		if err != nil {
			log.Fatal("receiveData(conn net.Conn): Error to conn.Read into jsonData: ", err)
		}

		if len(jsonDataTemp) > 0 {
			dataReceived = true
		}

	}

	/*var jsonData []byte
	copy(jsonData, jsonDataTemp)
	for i := 0; i < (len(jsonData)); i++ {
		if jsonData[i] == 0 {
			jsonData = jsonData[0 : len(jsonData)-1]

		}

	}*/

	fmt.Println("Data received ([] byte):\n ", jsonDataTemp)
	fmt.Println("Data received (string):\n ", string(jsonDataTemp))

	return jsonDataTemp
}
