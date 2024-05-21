package utils

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"net"
	"time"
)

// MakeTcpConnection opens a new TCP connection with the server at to net.Addr argument. Returns the connection if no
// errors were raised. Otherwise, it returns and logs errors.
func MakeTcpConnection(to net.Addr) (net.Conn, error) {
	duration, parseErr := time.ParseDuration("5s")

	if parseErr != nil {
		log.Error("duration parse error: %s", parseErr)
		return nil, parseErr
	}

	conn, connErr := net.DialTimeout("tcp", to.String(), duration)

	if connErr != nil {
		log.Error("error connecting to address via TCP: %s", connErr)
		return nil, connErr
	}

	return conn, nil
}
