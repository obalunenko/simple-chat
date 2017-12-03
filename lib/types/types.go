package types

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Client
type Client struct {
	address
	name string
	message
}

// address
type address struct {
	ip   string
	port string
}

// message
type message struct {
	timestamp   string
	messageText string
}

// SetAddress function set Port and IP for client
func (c *Client) SetAddress(ip string) {

	c.port = "8080"
	c.ip = ip

}

// GetAddress function return address of client in format IP:port
func (c *Client) GetAddress() (address string) {

	address = c.ip + ":" + c.port
	return address
}

// SetName function prompt to enter name of Client
func (c *Client) SetName() {
	fmt.Print("Enter your name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	nameInput, setNameErr := setNameReader.ReadString('\n')
	if setNameErr != nil {
		log.Fatal("Error: ", setNameErr)
	}
	nameInput = strings.Replace(nameInput, "\n", "", -1)
	c.name = nameInput

}

// Name say the name of Client
func (c *Client) Name() string {

	return c.name
}

// SetMessage function set message object:  messageText and Timestamp
func (c *Client) SetMessage() {

	c.SetMessageText()
	c.SetTimestamp()
}

// SetMessageText function prompt to enter text of message
func (c *Client) SetMessageText() {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	messageInput, readErr := reader.ReadString('\n')

	if readErr != nil {

		log.Fatal("Error: ", readErr)
	}

	c.messageText = messageInput

}

// SetTimestamp function set current timestamp for each message
func (c *Client) SetTimestamp() {

	var timestampLayout = "01-02-2006 15:46:02"
	t := time.Now()
	c.timestamp = t.Format(timestampLayout)
}

// GetMessage gives message with timestamp
func (c *Client) GetMessage() (message string) {

	message = c.messageText + "\t\t" + c.timestamp
	return message
}
