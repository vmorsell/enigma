package enigma

type Enigma struct {
	plugboard Plugboard
	spindle   Spindle
}

// New returns an Enigma instance.
func New(plugboard Plugboard, spindle Spindle) Enigma {
	return Enigma{
		plugboard: plugboard,
		spindle:   spindle,
	}
}

// Encrypt encrypts a single .
func (e Enigma) Encrypt(k Key) Key {
	k = e.plugboard.Handle(k)
	k = e.spindle.Handle(k)
	k = e.plugboard.Handle(k)
	return k
}
