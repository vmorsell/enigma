package enigma

type Rotor interface {
	Forward(k Key) Key
	Backward(k Key) Key
	Step()
	Position() Key
	Notch() Key
}

type rotor struct {
	forwardMap  map[Key]Key
	backwardMap map[Key]Key
	position    Key
	notch       Key
}

func NewRotor(forwardMap map[Key]Key, position, notch Key) Rotor {
	backwardMap := reverseMap(forwardMap)
	return &rotor{
		forwardMap,
		backwardMap,
		position,
		notch,
	}
}

func reverseMap(m map[Key]Key) map[Key]Key {
	res := make(map[Key]Key, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func (r *rotor) Forward(k Key) Key {
	k = k.Shift(r.position.Int())
	if v, ok := r.forwardMap[k]; ok {
		k = v
	}
	k = k.Shift(-r.position.Int())
	return k
}

func (r *rotor) Backward(k Key) Key {
	k = k.Shift(r.position.Int())
	if v, ok := r.backwardMap[k]; ok {
		k = v
	}
	k = k.Shift(-r.position.Int())
	return k
}

func (r *rotor) Step() {
	r.position--
	if r.position%26 == 0 {
		r.position = 0
	}
}

func (r *rotor) Position() Key {
	return r.position
}

func (r *rotor) Notch() Key {
	return r.notch
}
