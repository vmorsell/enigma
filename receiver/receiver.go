package main

import (
	"fmt"
	"log"
	"net"
)

type receiver struct {
	network    string
	port       int
	bufferSize int
	conn       net.PacketConn
}

func NewReceiver(network string, port int, bufferSize int) *receiver {
	return &receiver{
		network:    network,
		port:       port,
		bufferSize: bufferSize,
	}
}

type Message struct {
	addr    net.Addr
	payload []byte
}

func (r *receiver) Listen(msg chan Message) error {
	pc, err := net.ListenPacket(r.network, fmt.Sprintf(":%d", r.port))
	if err != nil {
		return fmt.Errorf("listen packet: %w", err)
	}
	defer pc.Close()

	buf := make([]byte, r.bufferSize)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Printf("err: %v", err)
		}
		msg <- Message{
			addr:    addr,
			payload: buf[:n],
		}
	}
}
