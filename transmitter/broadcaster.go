package main

import (
	"fmt"
	"net"
)

// broadcaster holds the logic for the broadcaster.
type broadcaster struct {
	network       string
	broadcastAddr string
	port          int
	conn          net.PacketConn
	dst           *net.UDPAddr
}

// NewBroadcaster returns a new broadcaster.
func NewBroadcaster(network, broadcastAddr string, port int) *broadcaster {
	return &broadcaster{
		network:       network,
		broadcastAddr: broadcastAddr,
		port:          port,
	}
}

// Connect initializes the broadcaster connection.
func (b *broadcaster) Connect() error {
	conn, err := net.ListenPacket(b.network, ":0")
	if err != nil {
		return fmt.Errorf("listen packet: %w", err)
	}

	dst, err := net.ResolveUDPAddr(b.network, fmt.Sprintf("%s:%d", b.broadcastAddr, b.port))
	if err != nil {
		return fmt.Errorf("resolve udp addr: %w", err)
	}

	b.conn = conn
	b.dst = dst
	return nil
}

// Broadcast transmist a single message.
func (b *broadcaster) Broadcast(msg string) error {
	if b.conn == nil {
		return fmt.Errorf("no connection available")
	}
	if b.dst == nil {
		return fmt.Errorf("no destination available")
	}

	_, err := b.conn.WriteTo([]byte(msg), b.dst)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

// Close terminates the broadcaster connection.
func (b *broadcaster) Close() error {
	if b.conn == nil {
		return fmt.Errorf("no connection to close")
	}
	b.conn.Close()
	return nil
}
