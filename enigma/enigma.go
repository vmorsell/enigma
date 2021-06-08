package enigma

type Enigma struct {
	Plugboard *Plugboard
}

// New returns an Enigma instance.
func New() *Enigma {
	return &Enigma{
		Plugboard: NewPlugboard(),
	}
}

// Encrypt encrypts a single .
func (e *Enigma) Encrypt(k Key) (Key, error) {
	if e.Plugboard != nil {
		k = e.Plugboard.Handle(k)
	}

	// todo(vm): add spindle

	if e.Plugboard != nil {
		k = e.Plugboard.Handle(k)
	}
	return k, nil
}
