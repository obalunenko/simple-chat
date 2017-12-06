package chatTypes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Client struct
type Client struct {
	address struct {
		// address nested struct that stores info about Ip and Port of Client
		ip   string
		port string
	}
	name    string
	message struct {
		// message nested struct that stores info about Timestamp and MessageText of Client's messages
		timestamp   string
		messageText string
	}
}

// IP function gives value of address.ip field of Client struct
func (c *Client) IP() string {
	return c.address.ip
}

// SetIp function sets value of address.ip field of Client struct
func (c *Client) setIP(ip string) {
	c.address.ip = ip
}

// Port function gives value of address.port field of Client struct
func (c *Client) Port() string {
	return c.address.port
}

// SetPort function sets value of address.port field of Client struct
func (c *Client) setPort(port string) {
	c.address.port = port
}

// SetAddress function set Port and IP for client
func (c *Client) SetAddress(ip string) {

	c.setPort("8080")
	c.setIP(ip)

}

// Address function return address of client in format IP:port
func (c *Client) Address() (address string) {

	address = c.address.ip + ":" + c.address.port
	return address
}

// SetName function prompt to enter name of Client
func (c *Client) SetName() {
	fmt.Print("Enter your name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	nameInput, setNameErr := setNameReader.ReadString('\n')
	if setNameErr != nil {
		log.Fatal("SetName(): Error to read input: ", setNameErr)
	}
	nameInput = strings.Replace(nameInput, "\n", "", -1)
	nameInput = strings.Replace(nameInput, "\r", "", -1)
	c.name = nameInput

}

// Name function gives  the name of Client struct
func (c *Client) Name() string {

	return c.name
}

// SetMessage function set message object:  messageText and Timestamp
func (c *Client) SetMessage() {

	c.setMessageText()
	c.setTimestamp()
}

// SetMessageText function prompt to enter text of message
func (c *Client) setMessageText() {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	messageInput, readErr := reader.ReadString('\n')

	if readErr != nil {

		log.Fatal("SetMessage(): Error to read input: ", readErr)
	}

	messageInput = strings.Replace(messageInput, "\n", "", -1)
	messageInput = strings.Replace(messageInput, "\r", "", -1)
	c.message.messageText = messageInput

}

// MessageText gives value of message.messageText field of Client struct
func (c *Client) MessageText() string {
	return c.message.messageText
}

// SetTimestamp function set current timestamp for each message
func (c *Client) setTimestamp() {

	t := time.Now()
	c.message.timestamp = t.Format("2006-01-02 15:04:05")

}

// Timestamp gives value of message.timestamp field of Client struct
func (c *Client) Timestamp() string {
	return c.message.timestamp
}

// Message gives message with timestamp and name of addressee
func (c *Client) Message() {

	message := c.Timestamp() + " - message from " + c.Name() + ": " + c.MessageText() + "\n"
	fmt.Println(message)
}

// ObjectToJson function creates json from client object
func (c *Client) ObjectToJson() (jsonData []byte) {

	var (
		clientJson = clientJsonType{
			Address: Address{
				IP:   c.Port(),
				Port: c.IP(),
			},
			Name: c.Name(),

			Message: Message{
				Timestamp:   c.Timestamp(),
				MessageText: c.MessageText(),
			},
		}
	)

	jsonData, jsonErr := json.Marshal(&clientJson)
	if jsonErr != nil {
		log.Fatal("ObjectToJson(): Error to Marshal: ", jsonErr)
	}

	return jsonData
}

// ObjectFromJson creates struct Client from json object
func (c *Client) ObjectFromJson(jsonData []byte) {

	clientJson := new(clientJsonType)

	err := json.Unmarshal(jsonData, &clientJson)
	if err != nil {
		log.Fatal("ObjectFromJson(): Error to Unmarshal: ", err)
	}

	c.name = clientJson.Name
	c.address.ip = clientJson.IP
	c.address.port = clientJson.Port
	c.message.messageText = clientJson.MessageText
	c.message.timestamp = clientJson.Timestamp

}
