package message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testMessage = Message{

	Name:      "TestName",
	Timestamp: "2017-12-07",
	Text:      "Test Message",
}

func TestClient_Name(t *testing.T) {
	Convey("#TestClient_Name() ", t, func() {
		Convey(" Should return Name of Message", func() {
			expectedResult := testMessage.Name

			result := testMessage.Name

			So(result, ShouldEqual, expectedResult)
		})
	})

}

func TestClient_NewClient(t *testing.T) {

	Convey("#TestClient_NewClient() ", t, func() {
		Convey("New Message should be created with passed parameters", func() {
			newTestClient := new(Message)
			newTestClient.NewClient(testMessage.Name,
				testMessage.Text,
				testMessage.Timestamp)

			So(newTestClient, ShouldResemble, &testMessage)
		})
	})

}

func TestClient_MessageText(t *testing.T) {

	Convey("#TestClient_MessageText() ", t, func() {
		Convey(" Should return Text of Message", func() {
			expectedResult := testMessage.Text

			result := testMessage.Text

			So(result, ShouldEqual, expectedResult)

		})
	})

}

func TestClient_ObjectFromJSON(t *testing.T) {

	Convey("#TestClient_ObjectFromJSON()", t, func() {
		Convey("When valid json got", func() {
			newTestClient := new(Message)
			jsonTestdata, err := ioutil.ReadFile("testdata/Valid_Client.json")
			if err != nil {
				fmt.Println("Error was occured during read from testdata/Valid_Client.json: ")

			}

			err = json.Unmarshal(jsonTestdata, &newTestClient)
			Convey("Should create Message object from valid json", func() {
				So(err, ShouldBeNil)
				So(newTestClient, ShouldResemble, &testMessage)

			})

		})
		Convey("When invalid json got", func() {
			newTestClient := new(Message)
			jsonTestdata, err := ioutil.ReadFile("testdata/Invalid_Client.json")
			jsonTestdata = append(jsonTestdata, 0) // add '\x00' to the end of valid file
			if err != nil {
				fmt.Println("Error was occurred during read from testdata/Invalid_Client.json: ")

			}

			err = json.Unmarshal(jsonTestdata, &newTestClient)
			Convey("Should throw error when invalid json got", func() {

				So(err, ShouldBeError)

			})

		})

	})
}

func TestClient_ObjectToJSON(t *testing.T) {
	Convey("#TestClient_ObjectToJSON()", t, func() {

		Convey("Should convert Message object to json byte array", func() {
			resultJSON, _ := json.Marshal(testMessage)

			expectedJSON, err := ioutil.ReadFile("testdata/expected.json")

			if err != nil {
				fmt.Println("Error was occurred during read from testdata/Invalid_Client.json: ")

			}
			So(resultJSON, ShouldResemble, expectedJSON)

		})

	})
}

func ExampleMessage() {

	testMessage.String()
	// OUTPUT:
	// 2017-12-07 - message from TestName: Test Message

}
