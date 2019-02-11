package tcp

import "net"

// Listener represents a server socket that can accept new client connections
type Listener struct {
	OnListen  chan string
	OnConnect chan *Client
	OnError   chan error
}

// Start runs the loop (and blocks forever) to accept new clients to the server.  This will automatically set up the
// client to asynchronously read and send packets using its stream.
func (obj *Listener) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			obj.OnError <- err
		}
	}()
	obj.OnListen <- address
	for {
		cxn, err := listener.Accept()
		if err != nil {
			obj.OnError <- err
		}
		client := &Client{
			Socket:       cxn,
			Stream:       make(chan []byte),
			OnDisconnect: make(chan error),
			IsClosed:     false,
		}
		go client.Receive()
		go client.Send()
		obj.OnConnect <- client
	}
}
