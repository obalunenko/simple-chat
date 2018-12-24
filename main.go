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

var ( // flags
	isHost = flag.Bool("listen", false, "Listens on the specified ip address")
	ip     = flag.String("ip", "", "server machine ip")
	port   = flag.String("port", "8080", "server port")
)

var (
	version string
	date    string
	commit  string
)

func main() {

	fmt.Printf("Version info: %s:%s\n", version, date)
	fmt.Printf("commit: %s\n", commit)

	flag.Parse()

	if *ip == "" {
		log.Fatalf("Server IP is not specified")
	}

	name, err := setName()
	if err != nil {
		log.Fatalf("Failed create client: %v", err)
	}

	cl := client.New(*isHost, *ip, *port, name)

	if err := cl.Run(); err != nil {
		log.Fatal(err)
	}

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
