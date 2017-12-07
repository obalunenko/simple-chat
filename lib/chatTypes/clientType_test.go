package tests

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/oleg-balunenko/simple-chat/lib/chatTypes"
)

func TestIP(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("#TestIP() ", func() {
		g.It(" Should return IP of Client", func() {
			testClient := new(chatTypes.Client)
			expectedResult := "100.100.100.100"

			testClient.NewClient(
				"TestName",
				"100.100.100.100",
				"8080",
				"Test message",
				"2017-12-07",
			)

			result := testClient.IP()
			g.Assert(result).Equal(expectedResult)

		})

	})

}

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

			result := testClient.Name()
			g.Assert(result).Equal(expectedResult)

		})
	})
}
