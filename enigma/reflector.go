package enigma

type Reflector interface {
	Reflect(k Key) Key
}

type reflector struct {
	mappings map[Key]Key
}

type reflectorConfig struct {
	mapping map[Key]Key
}

func NewReflector(config reflectorConfig) Reflector {
	return reflector{
		config.mapping,
	}
}

func (r reflector) Reflect(k Key) Key {
	if v, ok := r.mappings[k]; ok {
		return v
	}
	return k
}
