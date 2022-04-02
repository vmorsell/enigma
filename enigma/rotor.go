package enigma

type Rotor interface {
	Forward(k Key) Key
	Backward(k Key) Key
}

type rotor struct {
	forwardMap  map[Key]Key
	backwardMap map[Key]Key
}

func NewRotor(forwardMap map[Key]Key) Rotor {
	backwardMap := reverseMap(forwardMap)
	return rotor{
		forwardMap,
		backwardMap,
	}
}

func reverseMap(m map[Key]Key) map[Key]Key {
	res := make(map[Key]Key, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func (r rotor) Forward(k Key) Key {
	if v, ok := r.forwardMap[k]; ok {
		return v
	}
	return k
}

func (r rotor) Backward(k Key) Key {
	if v, ok := r.backwardMap[k]; ok {
		return v
	}
	return k
}
