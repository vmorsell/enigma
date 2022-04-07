package enigma

type Spindle interface {
	Handle(c Char) Char
	SetPositions(positions []Char)
}

type spindle struct {
	rotors    []Rotor
	reflector Reflector
}

func NewSpindle(rotorTypes []RotorType, reflectorType ReflectorType, rings []Char, positions []Char) Spindle {
	rot := make([]Rotor, 0, len(rotorTypes))
	for i := 0; i < len(rotorTypes); i++ {
		rot = append(rot, NewRotor(rotorTypes[i], rings[i], positions[i]))
	}
	ref := NewReflector(reflectorType)

	return &spindle{
		rot,
		ref,
	}
}

func (s *spindle) Handle(c Char) Char {
	rotate(s.rotors)

	// Forward-pass through rotors.
	for _, s := range s.rotors {
		c = s.Forward(c)
	}

	// Reflect.
	c = s.reflector.Reflect(c)

	// Backward-pass through rotors.
	for i := len(s.rotors) - 1; i >= 0; i-- {
		c = s.rotors[i].Backward(c)
	}

	return c
}

func (s *spindle) SetPositions(positions []Char) {
	for i, p := range positions {
		s.rotors[i].SetPosition(p)
	}
}

func rotate(rotors []Rotor) {
	for i := len(rotors) - 1; i >= 0; i-- {
		// First rotor is always incremented.
		if i == 0 {
			rotors[i].Step()
			continue
		}

		// A subsequent rotor is rotated if the previos
		// rotor is at its notch.
		if rotors[i-1].Position() == rotors[i-1].Notch() {
			rotors[i].Step()
			continue
		}
	}
}
