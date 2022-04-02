package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpindleHandle(t *testing.T) {
	rot := NewRotor(map[Key]Key{A: B, Y: X})
	ref := NewReflector(map[Key]Key{B: X})

	s := NewSpindle([]Rotor{rot}, ref)

	k := A
	want := Y

	res := s.Handle(k)
	require.Equal(t, want, res)
}
