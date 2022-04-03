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

	pb := NewPlugboard(map[Key]Key{
		A: M,
		F: I,
		N: V,
		P: S,
		T: U,
		W: Z,
	})
	e := New(pb, spindle)

	keys := []Key{S, E, C, R, E, T}
	want := []Key{L, C, G, O, D, U}

	got := e.Encrypt(keys)
	require.Equal(t, want, got)
}

func TestEncryptKey(t *testing.T) {
	rotors := []Rotor{
		NewRotor(RotorI, A, C),
		NewRotor(RotorII, B, B),
		NewRotor(RotorIII, C, A),
	}
	ref := NewReflector(ReflectorA)
	spindle := NewSpindle(rotors, ref)
	pb := NewPlugboard(map[Key]Key{
		A: X,
	})
	e := New(pb, spindle)

	in := A
	want := J

	res := e.encryptKey(in)
	require.Equal(t, want, res)
}
