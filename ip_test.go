package nethelper_test

import (
	"log"

	"github.com/northbright/nethelper"
)

func ExampleGetOutboundIP() {
	IP, err := nethelper.GetOutboundIP()
	if err != nil {
		log.Printf("GetOutboundIP failed. Error: %v", err)
		return
	}

	log.Printf("GetOutboundIP OK. IP: %s\n", IP.String())

	// Output:
}
