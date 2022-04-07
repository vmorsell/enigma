package enigma

type Plugboard interface {
	Handle(c Char) Char
}

type plugboard struct {
	forwardMap  map[Char]Char
	backwardMap map[Char]Char
}

type PlugboardMapping struct {
	From Char
	To   Char
}

func NewPlugboard(mappings []PlugboardMapping) Plugboard {
	forwardMap, backwardMap := plugboardMaps(mappings)
	return plugboard{
		forwardMap,
		backwardMap,
	}
}

func plugboardMaps(mapping []PlugboardMapping) (map[Char]Char, map[Char]Char) {
	forwardMap := make(map[Char]Char, len(mapping))
	backwardMap := make(map[Char]Char, len(mapping))
	for _, m := range mapping {
		forwardMap[m.From] = m.To
		backwardMap[m.To] = m.From
	}
	return forwardMap, backwardMap
}

func (p plugboard) Handle(c Char) Char {
	if v, ok := p.forwardMap[c]; ok {
		return v
	}
	if v, ok := p.backwardMap[c]; ok {
		return v
	}
	return c
}
