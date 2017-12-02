package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const port = "8080"

// Client
type Client struct {
	IP   string
	Name string
	Message
}

// SetName
func (c *Client) SetName() {
	fmt.Print("Enter your name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	name, setNameErr := setNameReader.ReadString('\n')
	if setNameErr != nil {
		log.Fatal("Error: ", setNameErr)
	}
	c.Name = name

}

// SetIP
func (c *Client) SetIP(ip string) {
	c.IP = ip
}

// Message
type Message struct {
	timestamp time.Time
	Message   string
}
