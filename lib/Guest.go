package lib

import "net"

// RunGuest takes an argument ip and connects to host with ip
func RunGuest(ip string) {

	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {

	}

}
