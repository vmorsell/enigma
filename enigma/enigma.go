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

// Encrypt encrypts a slice of chars.
func (e Enigma) Encrypt(chars []Char) []Char {
	res := make([]Char, 0, len(chars))
	for _, c := range chars {
		cc := e.encryptChar(c)
		res = append(res, cc)
	}
	return res
}

// encryptChar encrypts a single char.
func (e Enigma) encryptChar(c Char) Char {
	c = e.plugboard.Handle(c)
	c = e.spindle.Handle(c)
	c = e.plugboard.Handle(c)
	return c
}
