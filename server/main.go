package server

import (
	"fmt"

	"./net"
)

func LaunchServer(port int) {
	if err := net.StartServer(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
