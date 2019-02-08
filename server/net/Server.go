package net

import (
	"fmt"

	"./tcp"
)

func StartServer(address string) error {
	listener := tcp.Listener{
		OnListen:  make(chan string),
		OnConnect: make(chan *tcp.Client),
		OnError:   make(chan error),
	}
	go func() {
		for {
			select {
			case addr := <-listener.OnListen:
				fmt.Printf("Started server on %s.\n", addr)
				break
			case client := <-listener.OnConnect:
				client.Stream <- []byte("Hello, world!")
				break
			case err := <-listener.OnError:
				fmt.Println(err)
				break
			}
		}
	}()
	return listener.Start(address)
}
