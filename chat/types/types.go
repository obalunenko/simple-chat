package types

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Client struct
type (
	Client struct {
		Address `json:"address"`
		Name    string `json:"name"`
		Message `json:"message"`
	}
	Address struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	}
	Message struct {
		Timestamp string `json:"timestamp"`
		Text      string `json:"text"`
	}
)

// IP function gives value of Address.IP field of Client struct
func (c *Client) IP() string {
	return c.Address.IP
}

// SetIp function sets value of Address.IP field of Client struct
func (c *Client) setIP(ip string) {
	c.Address.IP = ip
}

// Port function gives value of Address.Port field of Client struct
func (c *Client) Port() string {
	return c.Address.Port
}

// SetPort function sets value of Address.Port field of Client struct
func (c *Client) setPort(port string) {
	c.Address.Port = port
}

// SetAddress function set Port and IP for guest
func (c *Client) SetAddress(ip string) (err error) {

	if ip != "" {
		c.setPort("8080")
		c.setIP(ip)
	} else {
		err = errors.New("error at SetAddress: IP could not be empty")
		return err
	}

	return err

}

// SetName function prompt to enter Name of Client
func (c *Client) SetName() (err error) {
	fmt.Print("Enter your Name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	nameInput, err := setNameReader.ReadString('\n')
	if err != nil {
		err = errors.New("SetName(): Error to read input: " + err.Error())
		return err
	}
	nameInput = strings.Replace(nameInput, "\n", "", -1)
	nameInput = strings.Replace(nameInput, "\r", "", -1)
	c.Name = nameInput

	return nil

}

// SetMessage function set message object:  Text and Timestamp
func (c *Client) SetMessage() (err error) {

	err = c.setMessageText()
	if err != nil {
		err = errors.New("SetMessage():  " + err.Error())
		return err
	}
	c.setTimestamp()

	return err
}

// SetMessageText function prompt to enter text of message
func (c *Client) setMessageText() (err error) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	messageInput, err := reader.ReadString('\n')

	if err != nil {

		err = errors.New("SetMessageText(): Error to read input: " + err.Error())
		return err
	}

	messageInput = strings.Replace(messageInput, "\n", "", -1)
	messageInput = strings.Replace(messageInput, "\r", "", -1)
	c.Message.Text = messageInput

	return nil

}

// MessageText gives value of message.Text field of Client struct
func (c *Client) MessageText() string {
	return c.Message.Text
}

// SetTimestamp function set current Timestamp for each message
func (c *Client) setTimestamp() {

	t := time.Now()
	c.Message.Timestamp = t.Format("2006-01-02 15:04:05")

}

// Timestamp gives value of message.Timestamp field of Client struct
func (c *Client) Timestamp() string {
	return c.Message.Timestamp
}

// Message gives message with Timestamp and Name of addressee
func (c *Client) MessageString() {

	message := c.Timestamp() + " - message from " + c.Name + ": " + c.MessageText() + "\n"
	fmt.Println(message)
}

// ObjectToJSON function creates json from guest object
func (c *Client) ObjectToJSON() (jsonData []byte, err error) {

	jsonData, jsonErr := json.Marshal(&c)
	if jsonErr != nil {
		jsonErr = errors.New("ObjectToJSON(): Error to Marshal: " + jsonErr.Error())
		return nil, jsonErr
	}

	return jsonData, nil
}

// ObjectFromJSON creates struct Client from json object
func (c *Client) ObjectFromJSON(jsonData []byte) error {

	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		err = errors.New("ObjectFromJSON(): Error to Unmarshal: " + err.Error())
		return err
	}

	return nil

}

func (c *Client) AddressString() string {
	return c.Address.IP + ":" + c.Address.Port
}

// NewClient creates new Client object with parameters from function arguments
func (c *Client) NewClient(name string, ip string, port string, messageText string, timestamp string) {
	c.Name = name
	c.Message.Text = messageText
	c.Message.Timestamp = timestamp
	c.Address.Port = port
	c.Address.IP = ip
}
