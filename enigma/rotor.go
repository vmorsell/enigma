package enigma

type Rotor interface {
	Forward(k Key) Key
	Backward(k Key) Key
	Step()
	SetRing(k Key)
	Position() Key
	SetPosition(k Key)
	Notch() Key
}

type rotor struct {
	forwardMap  map[Key]Key
	backwardMap map[Key]Key
	ring        Key
	position    Key
	notch       Key
}

type rotorConfig struct {
	mapping map[Key]Key
	notch   Key
}

func NewRotor(config rotorConfig, ring Key, position Key) Rotor {
	backwardMap := reverseMap(config.mapping)
	r := &rotor{
		forwardMap:  config.mapping,
		backwardMap: backwardMap,
		notch:       config.notch,
	}
	r.SetRing(ring)
	r.SetPosition(position)
	return r
}

func reverseMap(m map[Key]Key) map[Key]Key {
	res := make(map[Key]Key, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func (r *rotor) Forward(k Key) Key {
	k = k.Shift(r.position)
	if v, ok := r.forwardMap[k]; ok {
		k = v
	}
	k = k.Shift(-r.position)
	return k
}

func (r *rotor) Backward(k Key) Key {
	k = k.Shift(r.position)
	if v, ok := r.backwardMap[k]; ok {
		k = v
	}
	k = k.Shift(-r.position)
	return k
}

func (r *rotor) Step() {
	r.position = r.position.Shift(1)
}

// SetRing shifts the rotor mapping in the same way as shifting the wiring in
// the rotor would do.
func (r *rotor) SetRing(k Key) {
	r.position = r.position.Shift(-k)
	r.notch = r.notch.Shift(-k)
}

// SetPosition shifts the rotor mapping in the same way as a physical rotation
// of the rotor would do.
func (r *rotor) SetPosition(k Key) {
	r.position = r.position.Shift(k)
}

func (r *rotor) Position() Key {
	return r.position
}

func (r *rotor) Notch() Key {
	return r.notch
}
