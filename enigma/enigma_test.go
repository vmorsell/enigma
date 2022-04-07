package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	key := DailyKey{
		RotorTypes:    []RotorType{RotorIII, RotorI, RotorII},
		ReflectorType: ReflectorA,
		Rings:         []Char{V, M, X},
		Positions:     []Char{L, B, A},
		Plugs: []PlugboardMapping{
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
		RotorTypes:    []RotorType{RotorI, RotorII, RotorIII},
		ReflectorType: ReflectorA,
		Rings:         []Char{A, B, C},
		Positions:     []Char{C, B, A},
		Plugs: []PlugboardMapping{
			{A, X},
		},
	}
	e := NewEnigma(key)

	in := A
	want := J

	res := e.(*enigma).encryptChar(in)
	require.Equal(t, want, res)
}
