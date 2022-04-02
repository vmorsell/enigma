package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	in := A
	want := F

	rot := NewRotor(rotorConfig{
		mapping: map[Key]Key{
			B: C,
			E: D,
		},
		notch: Z,
	}, A)
	ref := NewReflector(map[Key]Key{C: D})

	s := NewSpindle([]Rotor{rot}, ref)
	p := NewPlugboard(map[Key]Key{A: B, F: E})

	e := New(p, s)
	res := e.Encrypt(in)
	require.Equal(t, want, res)
}
