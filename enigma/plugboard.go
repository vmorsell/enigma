package enigma

import "fmt"

type Plugboard struct {
	Patches map[Key]Key
}

func NewPlugboard() *Plugboard {
	return &Plugboard{
		Patches: make(map[Key]Key),
	}
}

// Handle processes a single character input through the Plugboard.
func (p *Plugboard) Handle(k Key) Key {
	if x, ok := p.Patches[k]; ok {
		return x
	}
	return k
}

// Patch connects two characters on the Patchboard. This makes the two
// characters switch place when processed by the board.
func (p *Plugboard) Patch(a, b Key) error {
	if v, ok := p.Patches[a]; ok {
		return fmt.Errorf("%s already patched to %s", a, v)
	}
	if v, ok := p.Patches[b]; ok {
		return fmt.Errorf("%s already patched to %s", b, v)
	}

	p.Patches[a] = b
	p.Patches[b] = a
	return nil
}
