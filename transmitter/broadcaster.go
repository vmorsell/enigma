package main

import (
	"fmt"
	"net"
)

type broadcaster struct {
	network       string
	broadcastAddr string
	port          int
	conn          net.PacketConn
	dst           *net.UDPAddr
}

func NewBroadcaster(network, broadcastAddr string, port int) *broadcaster {
	return &broadcaster{
		network:       network,
		broadcastAddr: broadcastAddr,
		port:          port,
	}
}

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

func (b *broadcaster) Close() error {
	if b.conn == nil {
		return fmt.Errorf("no connection to close")
	}
	b.conn.Close()
	return nil
}
