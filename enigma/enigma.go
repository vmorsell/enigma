package enigma

type Enigma struct {
	Plugboard *Plugboard
	spindle   Spindle
}

// New returns an Enigma instance.
func New(spindle Spindle) *Enigma {
	return &Enigma{
		Plugboard: NewPlugboard(),
		spindle:   spindle,
	}
}

// Encrypt encrypts a single .
func (e *Enigma) Encrypt(k Key) (Key, error) {
	if e.Plugboard != nil {
		k = e.Plugboard.Handle(k)
	}

	k = e.spindle.Handle(k)

	if e.Plugboard != nil {
		k = e.Plugboard.Handle(k)
	}
	return k, nil
}
