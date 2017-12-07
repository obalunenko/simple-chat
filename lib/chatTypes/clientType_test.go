<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> web-sockets
package chatTypes

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/franela/goblin"
)

var testClient = Client{
	address: struct {
		ip   string
		port string
	}{
		ip:   "100.100.100.100",
		port: "8080",
	},

	name: "TestName",

	message: struct {
		timestamp   string
		messageText string
	}{
		timestamp:   "2017-12-07",
		messageText: "Test Message",
	},
}

func TestClient_IP(t *testing.T) {
<<<<<<< HEAD
=======
package tests

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

func TestIP(t *testing.T) {
>>>>>>> 0a22fd0... lib/charTypes: Add tests for ClientType.go
=======
>>>>>>> web-sockets

	g := goblin.Goblin(t)

	g.Describe("#TestIP() ", func() {
		g.It(" Should return IP of Client", func() {
<<<<<<< HEAD
<<<<<<< HEAD
			expectedResult := testClient.address.ip
=======
			testClient := new(chatTypes.Client)
			expectedResult := "100.100.100.100"

			testClient.NewClient(
				"TestName",
				"100.100.100.100",
				"8080",
				"Test message",
				"2017-12-07",
			)
>>>>>>> 0a22fd0... lib/charTypes: Add tests for ClientType.go
=======
			expectedResult := testClient.address.ip
>>>>>>> web-sockets

			result := testClient.IP()
			g.Assert(result).Equal(expectedResult)

		})

	})

}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> web-sockets
func TestClient_Port(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestClient_Port() ", func() {
		g.It(" Should return Port of Client", func() {
			expectedResult := testClient.address.port

			result := testClient.Port()
			g.Assert(result).Equal(expectedResult)
		})
	})
}

func TestClient_Name(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestClient_Name() ", func() {
		g.It(" Should return Name of Client", func() {

			expectedResult := testClient.name
<<<<<<< HEAD
=======
func TestName(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestName() ", func() {
		g.It(" Should return Name of Client", func() {
			testClient := new(chatTypes.Client)
			expectedResult := "TestName"

			testClient.NewClient(
				"TestName",
				"100.100.100.100",
				"8080",
				"Test message",
				"2017-12-07",
			)
>>>>>>> 0a22fd0... lib/charTypes: Add tests for ClientType.go
=======
>>>>>>> web-sockets

			result := testClient.Name()
			g.Assert(result).Equal(expectedResult)

		})
	})
}
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> web-sockets

func TestClient_NewClient(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestClient_NewClient() ", func() {
		g.It("New Client should be created with passed parameters", func() {
			newTestClient := new(Client)
			newTestClient.NewClient(testClient.name,

				testClient.address.ip,
				testClient.address.port,
				testClient.message.messageText,
				testClient.message.timestamp)

			g.Assert(newTestClient).Equal(&testClient)
		})
	})

}

func TestClient_MessageText(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestClient_MessageText() ", func() {
		g.It(" Should return messageText of Client", func() {

			expectedResult := testClient.message.messageText

			result := testClient.MessageText()
			g.Assert(result).Equal(expectedResult)

		})
	})

}

func TestClient_ObjectFromJson(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestClient_ObjectFromJson()", func() {
		g.It(" Should create Client object from valid json", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testFiles/Valid_Client.json")
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Valid_Client.json: ")
				g.Fail(err)

			}
			newTestClient.ObjectFromJson(jsonTestdata)

			g.Assert(newTestClient).Equal(&testClient)

		})

		g.It(" Should throw error when invalid json got", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testFiles/Invalid_Client.json")
			jsonTestdata = append(jsonTestdata, 0)
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Invalid_Client.json: ")
				g.Fail(err)

			}

			var result bool // if error will be returned result will be true
			err = newTestClient.ObjectFromJson(jsonTestdata)
			if err != nil {

				result = true

			} else {
				result = false
			}

			expectedResult := true

			g.Assert(result).Equal(expectedResult)
		})

	})
}
<<<<<<< HEAD
=======
>>>>>>> 0a22fd0... lib/charTypes: Add tests for ClientType.go
=======
>>>>>>> web-sockets
