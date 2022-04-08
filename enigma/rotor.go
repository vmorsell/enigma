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
	typ      RotorType
	ring     Char
	position Char
}

type RotorType struct {
	name            string
	forwardMapping  map[Char]Char
	backwardMapping map[Char]Char
	notch           Char
}

func NewRotor(typ RotorType, ring Char, position Char) Rotor {
	r := &rotor{
		typ: typ,
	}
	r.SetPosition(position)
	r.SetRing(ring)
	return r
}

func (r *rotor) Forward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.typ.forwardMapping[c]; ok {
		c = v
	}
	c = c.Shift(-r.position)
	return c
}

func (r *rotor) Backward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.typ.backwardMapping[c]; ok {
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
	r.typ.notch = r.typ.notch.Shift(-c)
}

// SetPosition shifts the rotor mapping in the same way as a physical rotation
// of the rotor would do.
func (r *rotor) SetPosition(c Char) {
	r.position = c
}

func (r *rotor) Position() Char {
	return r.position
}

func (r *rotor) Notch() Char {
	return r.typ.notch
}
