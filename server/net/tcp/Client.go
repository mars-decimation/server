package tcp

import "net"

const (
	PacketLength int = 4096
)

type Client struct {
	Socket       net.Conn
	Stream       chan []byte
	OnDisconnect chan error
	IsClosed     bool
}

func (this *Client) Close() {
	this.IsClosed = true
	this.Socket.Close()
}

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

func (this *Client) Send() {
	for {
		select {
		case msg := <-this.Stream:
			this.Socket.Write(msg)
		}
	}
}
