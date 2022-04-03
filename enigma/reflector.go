package enigma

type Reflector interface {
	Reflect(c Char) Char
}

type reflector struct {
	mappings map[Char]Char
}

type reflectorConfig struct {
	mapping map[Char]Char
}

func NewReflector(config reflectorConfig) Reflector {
	return reflector{
		config.mapping,
	}
}

func (r reflector) Reflect(c Char) Char {
	if v, ok := r.mappings[c]; ok {
		return v
	}
	return c
}
