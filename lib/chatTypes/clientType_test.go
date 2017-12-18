package chatTypes

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

func TestClient_Port(t *testing.T) {

	Convey("#TestClient_Port()", t, func() {
		Convey("Should return Port of Client", func() {
			expectedResult := testClient.address.port

			result := testClient.Port()
			So(result, ShouldEqual, expectedResult)

		})
	})

}

func TestClient_IP(t *testing.T) {

	Convey("#TestClient_IP()", t, func() {
		Convey("Should return IP of Client", func() {
			expectedResult := testClient.address.ip

			result := testClient.IP()
			So(result, ShouldEqual, expectedResult)
		})
	})

}

func TestClient_SetAddress(t *testing.T) {
	Convey("#TestClient_SetAddress", t, func() {
		newTestClient := new(Client)
		Convey("When called with IP 192.192.100.100 as argument", func() {

			newTestClient.SetAddress("192.192.100.100")

			Convey("Port should be set up to 8080", func() {
				So(newTestClient.address.port, ShouldEqual, "8080")
			})
			Convey("IP should be set uo to 192.192.100.100", func() {
				So(newTestClient.address.ip, ShouldEqual, "192.192.100.100")
			})

		})
		Convey("When called with empty IP as argument", func() {

			Convey("Error should be returned", func() {
				So(newTestClient.SetAddress(""), ShouldBeError)
			})

		})

	})

}

func TestClient_Address(t *testing.T) {
	Convey("#TestClient_Address()", t, func() {
		Convey("When called should return fields of address nested struct for Client object", func() {

			result := testClient.Address()

			Convey("Address should equal to '100.100.100.100:8080'", func() {
				expectedResult := "100.100.100.100:8080"

				So(result, ShouldEqual, expectedResult)
			})
		})

	})
}

/*func TestClient_Message(t *testing.T) {
	Convey("#TestCLient_Message()", t, func() {
		Convey("When called, message should be composed using Client object fields", func() {
			result := testClient.Message()
		})
	})
}*/
func TestClient_Name(t *testing.T) {
	Convey("#TestClient_Name() ", t, func() {
		Convey(" Should return Name of Client", func() {
			expectedResult := testClient.name

			result := testClient.Name()

			So(result, ShouldEqual, expectedResult)
		})
	})

}

func TestClient_NewClient(t *testing.T) {

	Convey("#TestClient_NewClient() ", t, func() {
		Convey("New Client should be created with passed parameters", func() {
			newTestClient := new(Client)
			newTestClient.NewClient(testClient.name,

				testClient.address.ip,
				testClient.address.port,
				testClient.message.messageText,
				testClient.message.timestamp)

			So(newTestClient, ShouldResemble, &testClient)
		})
	})

}

func TestClient_MessageText(t *testing.T) {

	Convey("#TestClient_MessageText() ", t, func() {
		Convey(" Should return messageText of Client", func() {
			expectedResult := testClient.message.messageText

			result := testClient.MessageText()

			So(result, ShouldEqual, expectedResult)

		})
	})

}

func TestClient_ObjectFromJSON(t *testing.T) {

	Convey("#TestClient_ObjectFromJSON()", t, func() {
		Convey("When valid json got", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testFiles/Valid_Client.json")
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Valid_Client.json: ")

			}
			newTestClient.ObjectFromJSON(jsonTestdata)
			Convey("Should create Client object from valid json", func() {

				So(newTestClient, ShouldResemble, &testClient)

			})

		})
		Convey("When invalid json got", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testFiles/Invalid_Client.json")
			jsonTestdata = append(jsonTestdata, 0) // add '\x00' to the end of valid file
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Invalid_Client.json: ")

			}

			Convey("Should throw error when invalid json got", func() {

				So(newTestClient.ObjectFromJSON(jsonTestdata), ShouldBeError)

			})

		})

	})
}

func TestClient_ObjectToJSON(t *testing.T) {
	Convey("#TestClient_ObjectToJSON()", t, func() {

		Convey("Should convert Client object to json byte array", func() {
			resultJSON, _ := testClient.ObjectToJSON()

			expectedJSON, err := ioutil.ReadFile("testFiles/expected.json")

			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testFiles/Invalid_Client.json: ")

			}
			So(resultJSON, ShouldResemble, expectedJSON)

		})

	})
}

func ExampleClient_Message() {

	testClient.Message()
	// OUTPUT:
	// 2017-12-07 - message from TestName: Test Message

}
