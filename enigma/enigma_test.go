package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	key := DailyKey{
		Rotors:         []RotorType{RotorIII, RotorI, RotorII},
		Reflector:      ReflectorA,
		RingSettings:   []Char{V, M, X},
		RotorPositions: []Char{L, B, A},
		PlugConnections: []PlugboardMapping{
			{A, M},
			{F, I},
			{N, V},
			{P, S},
			{T, U},
			{W, Z},
		},
	}
	e := NewEnigma(key)

	in := []Char{S, E, C, R, E, T}
	want := []Char{L, C, G, O, D, U}

	got := e.Encrypt(in)
	require.Equal(t, want, got)
}

func TestEncryptChar(t *testing.T) {
	key := DailyKey{
		Rotors:         []RotorType{RotorI, RotorII, RotorIII},
		Reflector:      ReflectorA,
		RingSettings:   []Char{A, B, C},
		RotorPositions: []Char{C, B, A},
		PlugConnections: []PlugboardMapping{
			{A, X},
		},
	}
	e := NewEnigma(key)

	in := A
	want := J

	res := e.(*enigma).encryptChar(in)
	require.Equal(t, want, res)
}
