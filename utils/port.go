package utils

import (
	"fmt"
	"net"
)

// FindAvailablePort starts from the given port and increments until finding an open port
func FindAvailablePort(startPort int) int {
	port := startPort
	for {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			port++
			continue
		}
		listener.Close()
		return port
	}
}
