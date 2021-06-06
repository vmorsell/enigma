package enigma

import (
	"fmt"
	"strings"
)

type Enigma struct {
	Plugboard *Plugboard
}

// New returns an Enigma instance.
func New() *Enigma {
	return &Enigma{
		Plugboard: NewPlugboard(),
	}
}

// Encrypt encrypts a single character.
func (e *Enigma) Encrypt(v string) (string, error) {
	if len(v) != 1 {
		return "", fmt.Errorf("must encrypt one character a time")
	}

	v = strings.ToUpper(v)

	if e.Plugboard != nil {
		v = e.Plugboard.Handle(v)
	}

	// todo(vm): add spindle

	if e.Plugboard != nil {
		v = e.Plugboard.Handle(v)
	}
	return v, nil
}
