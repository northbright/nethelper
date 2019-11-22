package nethelper

import (
	"fmt"
	"net"
)

// GetOutboundIP gets the preferred outbound IP.
func GetOutboundIP() (net.IP, error) {
	remoteAddr := "8.8.8.8:80"

	// UDP, unlike TCP, it does not have handshake nor a connection.
	// It gets the local up address if it would connect to the remote addr.
	// The remote target does NOT need to be there.
	// It's a "logical" but NOT "physical" connection.
	// Follow the example on:
	// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	conn, err := net.Dial("udp", remoteAddr)
	if err != nil {
		return nil, fmt.Errorf("GetOutboundIP error: %v", err)
	}
	defer conn.Close()

	// Get UDP address from connection.
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
