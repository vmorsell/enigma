package enigma

type Reflector interface {
	Reflect(k Key) Key
}

type reflector struct {
	mappings map[Key]Key
}

func NewReflector(mappings map[Key]Key) Reflector {
	return reflector{
		mappings,
	}
}

func (r reflector) Reflect(k Key) Key {
	if v, ok := r.mappings[k]; ok {
		return v
	}
	return k
}
