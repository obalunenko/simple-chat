package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {

		log.Fatal("Error: ", dialErr)

	}

	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)

	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Fprint(conn, message)
}
