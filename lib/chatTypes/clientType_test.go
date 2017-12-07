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

	g := goblin.Goblin(t)

	g.Describe("#TestIP() ", func() {
		g.It(" Should return IP of Client", func() {
			expectedResult := testClient.address.ip

			result := testClient.IP()
			g.Assert(result).Equal(expectedResult)

		})

	})

}

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

			result := testClient.Name()
			g.Assert(result).Equal(expectedResult)

		})
	})
}

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
			jsonTestdata = append(jsonTestdata, 0) // add '\x00' to the end of valid file
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Invalid_Client.json: ")
				g.Fail(err)

			}

			var result bool // if error will be returned result will be true

			if newTestClient.ObjectFromJson(jsonTestdata) != nil {

				result = true

			} else {
				result = false
			}

			expectedResult := true

			g.Assert(result).Equal(expectedResult)

		})

	})
}
