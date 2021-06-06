package enigma

type Enigma struct {
	Plugboard *Plugboard
}

func New() *Enigma {
	return &Enigma{
		Plugboard: NewPlugboard(),
	}
}

func (e *Enigma) Encrypt(v string) string {
	if e.Plugboard != nil {
		v = e.Plugboard.Handle(v)
	}

	// todo(vm): add spindle

	if e.Plugboard != nil {
		v = e.Plugboard.Handle(v)
	}
	return v
}
