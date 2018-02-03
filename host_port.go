package netutil

import (
	"fmt"
	"net"
	"strconv"
)

// SplitHostPort split host and port and validate port to be an integer
// between 0 and 65535.
func SplitHostPort(hostPort string) (host string, port int, err error) {
	host, portStr, err := net.SplitHostPort(hostPort)
	if err != nil {
		return "", 0, fmt.Errorf("cannot split host and port, hostPort=%s, err=%v", hostPort, err)
	}
	port, err = strconv.Atoi(portStr)
	if err != nil || port < 0 || 65535 < port {
		return "", 0, fmt.Errorf("port must be integer between 0 and 65535, hostPort=%s, err=%v", hostPort, err)
	}
	return host, port, nil
}
