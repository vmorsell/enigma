package enigma

// Reflector holds the interface for a reflector.
type Reflector interface {
	Handle(c Char) Char
}

// reflector holds the reflector logic.
type reflector struct {
	typ ReflectorType
}

// ReflectorType represents the mapping of a specific type of reflector.
type ReflectorType struct {
	mapping map[Char]Char
}

// NewReflector returns a new reflector.
func NewReflector(typ ReflectorType) Reflector {
	return reflector{
		typ,
	}
}

// Reflect performs the substitution of the char.
func (r reflector) Handle(c Char) Char {
	if v, ok := r.typ.mapping[c]; ok {
		return v
	}
	return c
}
