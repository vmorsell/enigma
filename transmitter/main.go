package main

import (
	"log"
	"time"

	"github.com/vmorsell/enigma/enigma"
)

func main() {
	bc := NewBroadcaster("udp4", "255.255.255.255", 45092)
	if err := bc.Connect(); err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer bc.Close()

	dk := enigma.NewRandomDailyKey()
	e := enigma.NewEnigma(dk)
	for {
		time.Sleep(3 * time.Second)

		msg, err := randomMessage()
		if err != nil {
			log.Printf("random message: %v\n", err)
			continue
		}

		chars, err := enigma.StringToChars(msg)
		if err != nil {
			log.Printf("string to chars: %v\n", err)
			continue
		}

		mk := enigma.NewRandomMessageKey()
		digest := e.EncryptMessage(chars, dk, mk)

		if err := bc.Broadcast(digest.String()); err != nil {
			log.Printf("broadcast: %v", err)
		}

		log.Printf("Message key:        %v", mk.Positions)
		log.Printf("Message:            %s", msg)
		log.Printf("Digest transmitted: %s", digest.String())
	}
}
