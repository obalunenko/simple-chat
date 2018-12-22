package types

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testClient = Client{
	Address: Address{
		IP:   "100.100.100.100",
		Port: "8080",
	},

	Name: "TestName",

	Message: Message{
		Timestamp: "2017-12-07",
		Text:      "Test Message",
	},
}

func TestClient_Port(t *testing.T) {

	Convey("#TestClient_Port()", t, func() {
		Convey("Should return Port of Client", func() {
			expectedResult := testClient.Address.Port

			result := testClient.Port()
			So(result, ShouldEqual, expectedResult)

		})
	})

}

func TestClient_IP(t *testing.T) {

	Convey("#TestClient_IP()", t, func() {
		Convey("Should return IP of Client", func() {
			expectedResult := testClient.Address.IP

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
				So(newTestClient.Address.Port, ShouldEqual, "8080")
			})
			Convey("IP should be set uo to 192.192.100.100", func() {
				So(newTestClient.Address.IP, ShouldEqual, "192.192.100.100")
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
		Convey("When called should return fields of Address nested struct for Client object", func() {

			result := testClient.AddressString()

			Convey("Address should equal to '100.100.100.100:8080'", func() {
				expectedResult := "100.100.100.100:8080"

				So(result, ShouldEqual, expectedResult)
			})
		})

	})
}

func TestClient_Name(t *testing.T) {
	Convey("#TestClient_Name() ", t, func() {
		Convey(" Should return Name of Client", func() {
			expectedResult := testClient.Name

			result := testClient.Name

			So(result, ShouldEqual, expectedResult)
		})
	})

}

func TestClient_NewClient(t *testing.T) {

	Convey("#TestClient_NewClient() ", t, func() {
		Convey("New Client should be created with passed parameters", func() {
			newTestClient := new(Client)
			newTestClient.NewClient(testClient.Name,

				testClient.Address.IP,
				testClient.Address.Port,
				testClient.Message.Text,
				testClient.Message.Timestamp)

			So(newTestClient, ShouldResemble, &testClient)
		})
	})

}

func TestClient_MessageText(t *testing.T) {

	Convey("#TestClient_MessageText() ", t, func() {
		Convey(" Should return Text of Client", func() {
			expectedResult := testClient.Message.Text

			result := testClient.MessageText()

			So(result, ShouldEqual, expectedResult)

		})
	})

}

func TestClient_ObjectFromJSON(t *testing.T) {

	Convey("#TestClient_ObjectFromJSON()", t, func() {
		Convey("When valid json got", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testdata/Valid_Client.json")
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testdata/Valid_Client.json: ")

			}
			newTestClient.ObjectFromJSON(jsonTestdata)
			Convey("Should create Client object from valid json", func() {

				So(newTestClient, ShouldResemble, &testClient)

			})

		})
		Convey("When invalid json got", func() {
			newTestClient := new(Client)
			jsonTestdata, err := ioutil.ReadFile("testdata/Invalid_Client.json")
			jsonTestdata = append(jsonTestdata, 0) // add '\x00' to the end of valid file
			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testdata/Invalid_Client.json: ")

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

			expectedJSON, err := ioutil.ReadFile("testdata/expected.json")

			if err != nil {
				fmt.Println("Error was occured during read from lib/chatTypes/testdata/Invalid_Client.json: ")

			}
			So(resultJSON, ShouldResemble, expectedJSON)

		})

	})
}

func ExampleClient_Message() {

	testClient.MessageString()
	// OUTPUT:
	// 2017-12-07 - message from TestName: Test Message

}
