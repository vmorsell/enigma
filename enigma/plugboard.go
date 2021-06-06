package enigma

import "fmt"

type Plugboard struct {
	Patches map[string]string
}

func NewPlugboard() *Plugboard {
	return &Plugboard{
		Patches: make(map[string]string),
	}
}

func (p *Plugboard) Handle(v string) string {
	if x, ok := p.Patches[v]; ok {
		return x
	}
	return v
}

func (p *Plugboard) Patch(a, b string) error {
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
