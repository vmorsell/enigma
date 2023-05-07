package enigma

// Plugboard holds the interface for a plugboard.
type Plugboard interface {
	Handle(c Char) Char
}

// plugboard holds the plugboard logic.
type plugboard struct {
	mappings map[Char]Char
}

// PlugboardMapping represents a mapping between two characters on the plugboard.
type PlugboardMapping struct {
	From Char
	To   Char
}

// NewPlugboard returns a new plugboard instance.
func NewPlugboard(mappings []PlugboardMapping) Plugboard {
	m := substitutionMap(mappings)
	return plugboard{
		m,
	}
}

// substitutionMap compiles the forward and backward mapping for the plugboard.
func substitutionMap(mappings []PlugboardMapping) map[Char]Char {
	res := make(map[Char]Char, len(mappings)*2)
	for _, m := range mappings {
		res[m.From] = m.To
		res[m.To] = m.From
	}
	return res
}

// Handle substitutes a char based on the plugboard configuration.
func (p plugboard) Handle(c Char) Char {
	if v, ok := p.mappings[c]; ok {
		return v
	}
	return c
}
