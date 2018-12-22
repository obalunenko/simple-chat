package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/simple-chat/client"
)

var isHost bool
var ip string
var port string

func main() {

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.StringVar(&ip, "ip", "", "server machine ip")
	flag.StringVar(&port, "port", "8080", "server port")

	flag.Parse()

	if ip == "" {
		log.Fatalf("Server IP is not specified")
	}

	name, err := setName()
	if err != nil {
		log.Fatalf("Failed create client: %v", err)
	}

	cl := client.New(isHost, ip, port, name)

	cl.Run()

}

func setName() (string, error) {

	fmt.Print("Enter your Name: ")
	setNameReader := bufio.NewReader(os.Stdin)
	nameInput, err := setNameReader.ReadString('\n')
	if err != nil {
		err = errors.New("SetName(): Error to read input: " + err.Error())
		return "", err
	}
	nameInput = strings.Replace(nameInput, "\n", "", -1)
	nameInput = strings.Replace(nameInput, "\r", "", -1)
	return nameInput, nil
}
