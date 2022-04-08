package enigma

// Plugboard holds the interface for a plugboard.
type Plugboard interface {
	Handle(c Char) Char
}

// plugboard holds the plugboard logic.
type plugboard struct {
	forwardMap  map[Char]Char
	backwardMap map[Char]Char
}

// PlugboardMapping represents one connection between two chars
// in the plugboard.
type PlugboardMapping struct {
	From Char
	To   Char
}

// NewPlugboard returns a new plugboard instance.
func NewPlugboard(mappings []PlugboardMapping) Plugboard {
	forwardMap, backwardMap := plugboardMaps(mappings)
	return plugboard{
		forwardMap,
		backwardMap,
	}
}

// plugboardMaps returns the forward and backward mapping for the plugboard.
func plugboardMaps(mapping []PlugboardMapping) (map[Char]Char, map[Char]Char) {
	forwardMap := make(map[Char]Char, len(mapping))
	backwardMap := make(map[Char]Char, len(mapping))
	for _, m := range mapping {
		forwardMap[m.From] = m.To
		backwardMap[m.To] = m.From
	}
	return forwardMap, backwardMap
}

// Handle substitutes a char based on the plugboard configuration.
func (p plugboard) Handle(c Char) Char {
	if v, ok := p.forwardMap[c]; ok {
		return v
	}
	if v, ok := p.backwardMap[c]; ok {
		return v
	}
	return c
}
