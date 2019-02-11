package tcp

import "net"

const (
	// PacketLength is the maximum length of each packet that should be received as one packet
	PacketLength int = 4096
)

// Client describes a socket that has connected to this server
type Client struct {
	Socket       net.Conn
	Stream       chan []byte
	OnDisconnect chan error
	IsClosed     bool
}

// Close disconnects the client
func (client *Client) Close() {
	client.IsClosed = true
	client.Socket.Close()
}

// Receive runs a loop (and blocks forever) to read packets from the socket
func (client *Client) Receive() {
	for {
		message := make([]byte, PacketLength)
		len, err := client.Socket.Read(message)
		if err != nil {
			if !client.IsClosed {
				client.OnDisconnect <- err
				client.Close()
			}
			break
		}
		if len > 0 {
			client.Stream <- message[0:len]
		}
	}
}

// Send runs a loop (and blocks forever) to send packets to the socket
func (client *Client) Send() {
	for {
		select {
		case msg := <-client.Stream:
			client.Socket.Write(msg)
		}
	}
}
