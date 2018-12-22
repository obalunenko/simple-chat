package message

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
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
		return errors.Wrap(err, "message: SetMessage")
	}

	m.Name = from
	m.Text = text
	m.setTimestamp()

	return nil
}

// inputMessageText function prompt to enter text of message
func inputMessageText() (text string, err error) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	messageInput, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.Wrap(err, "message: inputMessageText")
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
func (m *Message) Send(conn io.Writer) error {

	jsonData, err := json.Marshal(&m)
	if err != nil {
		return errors.Wrap(err, "message: Send")
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		return errors.Wrap(err, "message: Send")
	}

	return nil
}

// Receive receives byte array from connection net.Conn
func (m *Message) Receive(conn io.Reader) error {

	dataReceived := false

	jsonData := make([]byte, 1000)

	for !dataReceived {
		_, err := conn.Read(jsonData)
		if err != nil {
			return errors.Wrap(err, "message: Receive")
		}

		if len(jsonData) > 0 {
			dataReceived = true
		}

	}

	jsonData = bytes.Trim(jsonData, "\x00")

	err := json.Unmarshal(jsonData, &m)
	if err != nil {
		return errors.Wrap(err, "message: Receive")

	}

	return nil
}
