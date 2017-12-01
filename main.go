package main

import (
	"flag"
	"log"

	"github.com/oleg-balunenko/simple-chat/lib"
)

//noinspection GoUnresolvedReference
func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Error:ip address not specified")
	}
	connIP := flag.Args()[0]

	if isHost {
		// go run  main.go  -listen <ip>

		lib.RunHost(connIP)

	} else {
		// go run main.go  <ip>

		lib.RunGuest(connIP)
	}

}
