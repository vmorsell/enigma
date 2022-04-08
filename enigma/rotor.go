package enigma

// Rotor holds the interface for a rotor.
type Rotor interface {
	Forward(c Char) Char
	Backward(c Char) Char
	Step()
	SetRing(c Char)
	Position() Char
	SetPosition(c Char)
	Notch() Char
}

// rotor holds the rotor logic.
type rotor struct {
	typ      RotorType
	ring     Char
	position Char
}

// RotorType is the model for a specific rotor type.
type RotorType struct {
	name            string
	forwardMapping  map[Char]Char
	backwardMapping map[Char]Char
	notch           Char
}

// NewRotor returns a new rotor.
func NewRotor(typ RotorType, ring Char, position Char) Rotor {
	r := &rotor{
		typ: typ,
	}
	r.SetPosition(position)
	r.SetRing(ring)
	return r
}

// Forward performs the forward substitution of a char through the rotor.
func (r *rotor) Forward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.typ.forwardMapping[c]; ok {
		c = v
	}
	c = c.Shift(-r.position)
	return c
}

// Backward performs the backward substitution of a char through the rotor.
func (r *rotor) Backward(c Char) Char {
	c = c.Shift(r.position)
	if v, ok := r.typ.backwardMapping[c]; ok {
		c = v
	}
	c = c.Shift(-r.position)
	return c
}

// Step rotates the rotor one step.
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

// Position returns the current position of the rotor.
func (r *rotor) Position() Char {
	return r.position
}

// Notch returns the current notch position of the rotor.
func (r *rotor) Notch() Char {
	return r.typ.notch
}
