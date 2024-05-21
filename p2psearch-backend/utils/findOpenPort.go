package utils

import (
	"fmt"
	"net"
)

// FindOpenPort finds an open port in the range [startPort, endPort]
func FindOpenPort(startPort, endPort int) (string, error) {
	for port := startPort; port <= endPort; port++ {
		// Attempt to bind to this port
		listener, listenerErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if listenerErr == nil {
			// If we were able to bind, close the listener and return the port
			listener.Close()
			return fmt.Sprintf(":%d", port), nil
		}
	}
	// If we were unable to bind to any ports, return an empty string
	return "", fmt.Errorf("unable to find open port")
}
