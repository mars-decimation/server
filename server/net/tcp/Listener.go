package tcp

import "net"

type Listener struct {
	OnListen  chan string
	OnConnect chan *Client
	OnError   chan error
}

func (this *Listener) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			this.OnError <- err
		}
	}()
	this.OnListen <- address
	for {
		cxn, err := listener.Accept()
		if err != nil {
			this.OnError <- err
		}
		client := &Client{
			Socket:       cxn,
			Stream:       make(chan []byte),
			OnDisconnect: make(chan error),
			IsClosed:     false,
		}
		go client.Receive()
		go client.Send()
		this.OnConnect <- client
	}
}
