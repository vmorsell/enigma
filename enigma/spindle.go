package enigma

type Spindle interface {
	Handle(k Key) Key
}

type spindle struct {
	rotors    []Rotor
	reflector Reflector
}

func NewSpindle(rotors []Rotor, reflector Reflector) Spindle {
	return spindle{
		rotors,
		reflector,
	}
}

func (s spindle) Handle(k Key) Key {
	// Forward-pass through rotors.
	for _, s := range s.rotors {
		k = s.Forward(k)
	}

	// Reflect.
	k = s.reflector.Reflect(k)

	// Backward-pass through rotors.
	for i := len(s.rotors) - 1; i >= 0; i-- {
		k = s.rotors[i].Backward(k)
	}

	// Step the rotors.
	for i := len(s.rotors) - 1; i >= 0; i-- {
		// First rotor is always incremented.
		if i == 0 {
			s.rotors[i].Step()
			continue
		}

		// A subsequent rotor is rotated if the previos
		// rotor is at its notch.
		if s.rotors[i-1].Position() == s.rotors[i-1].Notch() {
			s.rotors[i].Step()
			continue
		}
	}

	return k
}
