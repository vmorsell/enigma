package enigma

// Spindle holds the interface for a spindle.
type Spindle interface {
	Handle(c Char) Char
	SetPositions(positions []Char)
}

// spindle holds the spindle logic.
type spindle struct {
	rotors    []Rotor
	reflector Reflector
}

// NewSpindle returns a new spindle.
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

// Handle substitutes a char in a forward-backward pass through the spindle.
func (s *spindle) Handle(c Char) Char {
	rotate(s.rotors)

	// Forward-pass through rotors.
	for _, s := range s.rotors {
		c = s.Forward(c)
	}

	// Reflect.
	c = s.reflector.Handle(c)

	// Backward-pass through rotors.
	for i := len(s.rotors) - 1; i >= 0; i-- {
		c = s.rotors[i].Backward(c)
	}

	return c
}

// SetPositions updates the positions of the rotors in the spindle.
func (s *spindle) SetPositions(positions []Char) {
	for i, p := range positions {
		s.rotors[i].SetPosition(p)
	}
}

// rotate performs one step shift of the appropriate rotors. This is done every
// time a key is pressed on the Enigma machine.
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
