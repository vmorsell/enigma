package enigma

type Reflector interface {
	Reflect(c Char) Char
}

type reflector struct {
	typ ReflectorType
}

type ReflectorType struct {
	mapping map[Char]Char
}

func NewReflector(typ ReflectorType) Reflector {
	return reflector{
		typ,
	}
}

func (r reflector) Reflect(c Char) Char {
	if v, ok := r.typ.mapping[c]; ok {
		return v
	}
	return c
}
