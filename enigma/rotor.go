package enigma

type Rotor interface {
	Forward(c Char) Char
	Backward(c Char) Char
	Step()
	SetRing(c Char)
	Position() Char
	SetPosition(c Char)
	Notch() Char
}

type rotor struct {
	forwardMap  map[Char]Char
	backwardMap map[Char]Char
	ring        Char
	position    Char
	notch       Char
}

type rotorConfig struct {
	mapping map[Char]Char
	notch   Char
}

func NewRotor(config rotorConfig, ring Char, position Char) Rotor {
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

func reverseMap(m map[Char]Char) map[Char]Char {
	res := make(map[Char]Char, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func (r *rotor) Forward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.forwardMap[c]; ok {
		c = v
	}
	c = c.Shift(-r.position)
	return c
}

func (r *rotor) Backward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.backwardMap[c]; ok {
		c = v
	}
	c = c.Shift(-r.position)
	return c
}

func (r *rotor) Step() {
	r.position = r.position.Shift(1)
}

// SetRing shifts the rotor mapping in the same way as shifting the wiring in
// the rotor would do.
func (r *rotor) SetRing(c Char) {
	r.position = r.position.Shift(-c)
	r.notch = r.notch.Shift(-c)
}

// SetPosition shifts the rotor mapping in the same way as a physical rotation
// of the rotor would do.
func (r *rotor) SetPosition(c Char) {
	r.position = r.position.Shift(c)
}

func (r *rotor) Position() Char {
	return r.position
}

func (r *rotor) Notch() Char {
	return r.notch
}
