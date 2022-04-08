package main

import "log"

func main() {
	r := NewReceiver("udp4", 45092, 1024)

	msgs := make(chan Message)
	go r.Listen(msgs)

	for {
		msg := <-msgs
		log.Printf("intercepted: %s", msg.payload)
	}
}
