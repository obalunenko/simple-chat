package message

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const prefix = "message"

// Message struct
type (
	Message struct {
		Name      string `json:"name"`
		Timestamp string `json:"timestamp"`
		Text      string `json:"text"`
	}
)

// SetMessage function set message object:  Text and Timestamp
func (m *Message) SetMessage(from string, reader io.Reader) (err error) {
	var funcName = "SetMessage"
	text, err := inputMessageText(reader)

	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	m.Name = from
	m.Text = text
	m.setTimestamp()

	return nil
}

// inputMessageText function prompt to enter text of message
func inputMessageText(reader io.Reader) (text string, err error) {
	var funcName = "inputMessageText()"
	fmt.Print("Send message: ")
	r := bufio.NewReader(reader)

	messageInput, err := r.ReadString('\n')

	if err != nil {
		return "", errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
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
func (m *Message) String() string {

	return fmt.Sprintf("%s - message from %s: %s\n", m.Timestamp, m.Name, m.Text)

}

// New creates new Message object with parameters from function arguments
func (m *Message) New(name string, messageText string, timestamp string) {

	m.Name = name
	m.Text = messageText
	m.Timestamp = timestamp

}

// Send sends byte array of Message to connection net.Conn
func (m *Message) Send(conn io.Writer) error {
	var funcName = "Send()"

	jsonData, err := json.Marshal(&m)
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
	}

	return nil
}

// Receive receives byte array from connection net.Conn
func (m *Message) Receive(conn io.Reader) error {
	var funcName = "Receive()"
	dataReceived := false

	jsonData := make([]byte, 1000)

	for !dataReceived {
		_, err := conn.Read(jsonData)
		if err != nil {
			return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))
		}

		if len(jsonData) > 0 {
			dataReceived = true
		}

	}

	jsonData = bytes.Trim(jsonData, "\x00")

	err := json.Unmarshal(jsonData, &m)
	if err != nil {
		return errors.Wrap(err, strings.Join([]string{prefix, funcName}, ":"))

	}

	return nil
}
