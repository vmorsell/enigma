package enigma

// Enigma represents an Enigma instance.
type Enigma interface {
	Encrypt(chars []Char) []Char
	SetDailyKey(key DailyKey)
	SetMessageKey(key MessageKey)
}

// enigma implements the Enigma logic.
type enigma struct {
	plugboard Plugboard
	spindle   Spindle
}

// New returns an Enigma instance.
func NewEnigma(key DailyKey) Enigma {
	e := &enigma{}
	e.SetDailyKey(key)
	return e
}

// SetDailyKey applies a daily key for the Enigma instance.
func (e *enigma) SetDailyKey(key DailyKey) {
	pb := NewPlugboard(key.Plugs)
	spindle := NewSpindle(key.RotorTypes, key.ReflectorType, key.Rings, key.Positions)

	e.plugboard = pb
	e.spindle = spindle
}

// SetMessageKey applies a message key to the Enigma instance.
func (e *enigma) SetMessageKey(key MessageKey) {
}

// Encrypt encrypts a slice of chars.
func (e *enigma) Encrypt(chars []Char) []Char {
	res := make([]Char, 0, len(chars))
	for _, c := range chars {
		cc := e.encryptChar(c)
		res = append(res, cc)
	}
	return res
}

// encryptChar encrypts a single char.
func (e *enigma) encryptChar(c Char) Char {
	c = e.plugboard.Handle(c)
	c = e.spindle.Handle(c)
	c = e.plugboard.Handle(c)
	return c
}
