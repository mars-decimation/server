package server

import (
	"fmt"

	"./net"
)

// LaunchServer launches all of the server code and blocks forever
func LaunchServer(port int) {
	if err := net.StartServer(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
