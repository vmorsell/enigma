package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	rot := []Rotor{
		NewRotor(RotorIII, V, L),
		NewRotor(RotorI, M, B),
		NewRotor(RotorII, X, A),
	}
	ref := NewReflector(ReflectorA)
	spindle := NewSpindle(rot, ref)

	pb := NewPlugboard(PlugboardSettings{
		Mappings: []PlugboardMapping{
			{A, M},
			{F, I},
			{N, V},
			{P, S},
			{T, U},
			{W, Z},
		},
	})
	e := New(pb, spindle)

	chars := []Char{S, E, C, R, E, T}
	want := []Char{L, C, G, O, D, U}

	got := e.Encrypt(chars)
	require.Equal(t, want, got)
}

func TestEncryptChar(t *testing.T) {
	rotors := []Rotor{
		NewRotor(RotorI, A, C),
		NewRotor(RotorII, B, B),
		NewRotor(RotorIII, C, A),
	}
	ref := NewReflector(ReflectorA)
	spindle := NewSpindle(rotors, ref)
	pb := NewPlugboard(PlugboardSettings{
		Mappings: []PlugboardMapping{
			{A, X},
		},
	})
	e := New(pb, spindle)

	in := A
	want := J

	res := e.encryptChar(in)
	require.Equal(t, want, res)
}
