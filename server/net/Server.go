package net

import (
	"fmt"

	"../buildconfig"
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
				go func() {
					client.Stream <- []byte(fmt.Sprintf("%s (%s)\nOfficial Server\nOK\n", buildconfig.Config.Product, buildconfig.Config.Version))
					DoLogin(client)
				}()
				break
			case err := <-listener.OnError:
				fmt.Println(err)
				break
			}
		}
	}()
	return listener.Start(address)
}
