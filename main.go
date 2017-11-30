package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oleg-balunenko/simple-chat/lib"
)

//noinspection GoUnresolvedReference
func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if len(os.Args) <= 1 {
		log.Fatal("Error: ip address not specified")
	}

	fmt.Println("Length of arguments: ", len(os.Args))

	if isHost {
		// go run  main.go  -listen <ip>
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		// go run main.go  <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}

}
