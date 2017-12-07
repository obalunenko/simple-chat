package lib

import (
	"bytes"
	"errors"
	"log"
	"net"

	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

// sendData sends byte array of Client to connection net.Conn
func sendData(c *chatTypes.Client, conn net.Conn) (err error) {

	jsonDataToSend := make([]byte, 500)
	jsonDataToSend, err = c.ObjectToJson()
	if err != nil {
		err = errors.New("Error at sendData(c Client) at ObjectToJson(): " + err.Error())
		return err

	}

	_, err = conn.Write(jsonDataToSend)
	if err != nil {
		err = errors.New("Error at sendData(c Client) at conn.Write(): " + err.Error())
		return err
	}

	return nil
}

// receiveData receives byte array from connection net.Conn
func receiveData(conn net.Conn) []byte {

	dataReceived := false

	jsonData := make([]byte, 1000)

	for !dataReceived {
		_, err := conn.Read(jsonData)
		if err != nil {
			log.Fatal("receiveData(conn net.Conn): Error to conn.Read into jsonData: ", err)
		}

		if len(jsonData) > 0 {
			dataReceived = true
		}

	}

	jsonData = bytes.Trim(jsonData, "\x00")

	return jsonData
}
