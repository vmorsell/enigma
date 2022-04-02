package enigma

type Plugboard interface {
	Forward(k Key) Key
	Backward(k Key) Key
}

type plugboard struct {
	forwardMap  map[Key]Key
	backwardMap map[Key]Key
}

func NewPlugboard(forwardMap map[Key]Key) Plugboard {
	backwardMap := reverseMap(forwardMap)
	return plugboard{
		forwardMap,
		backwardMap,
	}
}

func (p plugboard) Forward(k Key) Key {
	if v, ok := p.forwardMap[k]; ok {
		return v
	}
	return k
}

func (p plugboard) Backward(k Key) Key {
	if v, ok := p.backwardMap[k]; ok {
		return v
	}
	return k
}
