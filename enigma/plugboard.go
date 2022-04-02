package enigma

type Plugboard interface {
	Handle(k Key) Key
}

type plugboard struct {
	mapping map[Key]Key
}

func NewPlugboard(mapping map[Key]Key) Plugboard {
	return plugboard{
		mapping,
	}
}

func (p plugboard) Handle(key Key) Key {
	for k, v := range p.mapping {
		if k == key {
			return v
		}
		if v == key {
			return k
		}
	}
	return key
}
