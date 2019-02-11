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
func (this *Client) Close() {
	this.IsClosed = true
	this.Socket.Close()
}

// Receive runs a loop (and blocks forever) to read packets from the socket
func (this *Client) Receive() {
	for {
		message := make([]byte, PacketLength)
		len, err := this.Socket.Read(message)
		if err != nil {
			if !this.IsClosed {
				this.OnDisconnect <- err
				this.Close()
			}
			break
		}
		if len > 0 {
			this.Stream <- message[0:len]
		}
	}
}

// Send runs a loop (and blocks forever) to send packets to the socket
func (this *Client) Send() {
	for {
		select {
		case msg := <-this.Stream:
			this.Socket.Write(msg)
		}
	}
}
