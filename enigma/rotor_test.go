package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseMap(t *testing.T) {
	in := map[Key]Key{A: B}
	want := map[Key]Key{B: A}

	res := reverseMap(in)
	require.Equal(t, want, res)
}

func TestForwardAndBackward(t *testing.T) {
	conf := rotorConfig{
		mapping: map[Key]Key{
			A: B,
		},
		notch: Z,
	}
	r := NewRotor(conf, A)

	t.Run("forward", func(t *testing.T) {

		tests := []struct {
			name string
			k    Key
			res  Key
		}{
			{
				name: "key found",
				k:    A,
				res:  B,
			},
			{
				name: "key not found",
				k:    X,
				res:  X,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := r.Forward(tt.k)
				require.Equal(t, tt.res, res)
			})
		}
	})

	t.Run("backward", func(t *testing.T) {

		tests := []struct {
			name string
			k    Key
			res  Key
		}{
			{
				name: "key found",
				k:    B,
				res:  A,
			},
			{
				name: "key not found",
				k:    X,
				res:  X,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := r.Backward(tt.k)
				require.Equal(t, tt.res, res)
			})
		}
	})
}

func TestStep(t *testing.T) {
	pos := B
	want := A
	r := NewRotor(rotorConfig{}, pos)

	r.Step()
	require.Equal(t, r.Position(), want)
}
