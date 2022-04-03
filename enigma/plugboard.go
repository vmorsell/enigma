package enigma

type Plugboard interface {
	Handle(c Char) Char
}

type plugboard struct {
	mapping map[Char]Char
}

func NewPlugboard(mapping map[Char]Char) Plugboard {
	return plugboard{
		mapping,
	}
}

func (p plugboard) Handle(c Char) Char {
	for k, v := range p.mapping {
		if k == c {
			return v
		}
		if v == c {
			return k
		}
	}
	return c
}
