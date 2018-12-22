package message

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Message struct
type (
	Message struct {
		Name      string `json:"name"`
		Timestamp string `json:"timestamp"`
		Text      string `json:"text"`
	}
)

// SetMessage function set message object:  Text and Timestamp
func (m *Message) SetMessage(from string) (err error) {

	text, err := inputMessageText()

	if err != nil {
		err = errors.New("SetMessage():  " + err.Error())
		return err
	}

	m.Name = from
	m.Text = text
	m.setTimestamp()

	return err
}

// inputMessageText function prompt to enter text of message
func inputMessageText() (text string, err error) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	messageInput, err := reader.ReadString('\n')

	if err != nil {

		err = errors.New("SetMessageText(): Error to read input: " + err.Error())
		return "", err
	}

	messageInput = strings.Replace(messageInput, "\n", "", -1)
	messageInput = strings.Replace(messageInput, "\r", "", -1)

	return messageInput, nil

}

// setTimestamp function set current Timestamp for each message
func (m *Message) setTimestamp() {

	t := time.Now()
	m.Timestamp = t.Format("2006-01-02 15:04:05")

}

// String gives message with Timestamp and Name of addressee
func (m *Message) String() {

	message := m.Timestamp + " - message from " + m.Name + ": " + m.Text + "\n"
	fmt.Println(message)
}

// NewClient creates new Message object with parameters from function arguments
func (m *Message) NewClient(name string, messageText string, timestamp string) {

	m.Name = name
	m.Text = messageText
	m.Timestamp = timestamp

}

// Send sends byte array of Message to connection net.Conn
func (m *Message) Send(conn io.Writer) (err error) {

	jsonData, jsonErr := json.Marshal(&m)
	if jsonErr != nil {
		jsonErr = errors.New("failed to to Marshal: " + jsonErr.Error())
		return jsonErr
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		err = errors.New("Error at Send at conn.Write(): " + err.Error())
		return err
	}

	return nil
}

// Receive receives byte array from connection net.Conn
func (m *Message) Receive(conn io.Reader) {

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

	err := json.Unmarshal(jsonData, &m)
	if err != nil {

		log.Fatalf("Failed to Unmarshal: %v", err)

	}

}
