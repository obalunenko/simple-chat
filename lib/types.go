package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const port = "8080"

// Client
type Client struct {
	IP   string
	Name string
	Message
}

// Message
type Message struct {
	Timestamp   time.Time
	MessageText string
}

// SetName
func (c *Client) SetName() {
	fmt.Print("Enter your name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	name, setNameErr := setNameReader.ReadString('\n')
	if setNameErr != nil {
		log.Fatal("Error: ", setNameErr)
	}
	name = strings.Replace(name, "\n", "", -1)
	c.Name = name

}

// SetIP
func (c *Client) SetIP(ip string) {
	c.IP = ip
}

// SetMessageText
func (c *Client) SetMessageText() {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	c.MessageText = message
}

// SetTimestamp
func (c *Client) SetTimestamp() {

	c.Timestamp = time.Now()

}
