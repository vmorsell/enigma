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

	rotate(s.rotors)

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

	return k
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
