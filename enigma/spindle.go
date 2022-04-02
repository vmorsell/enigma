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

	return k
}
