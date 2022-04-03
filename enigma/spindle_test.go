package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpindleHandle(t *testing.T) {
	tests := []struct {
		name  string
		rot   []Rotor
		ref   Reflector
		chars []Char
		res   []Char
	}{
		{
			name:  "one rotor - one char",
			rot:   []Rotor{NewRotor(RotorI, A, A)},
			ref:   NewReflector(ReflectorA),
			chars: []Char{A},
			res:   []Char{X},
		},
		{
			name:  "one rotor - multiple chars",
			rot:   []Rotor{NewRotor(RotorI, A, A)},
			ref:   NewReflector(ReflectorA),
			chars: []Char{A, A},
			res:   []Char{X, M},
		},
		{
			name: "two rotors - multiple chars",
			rot: []Rotor{
				NewRotor(RotorI, A, A),
				NewRotor(RotorII, A, A),
			},
			ref:   NewReflector(ReflectorA),
			chars: []Char{A, A},
			res:   []Char{X, M},
		},
		{
			name: "three rotors",
			rot: []Rotor{
				NewRotor(RotorIII, A, A),
				NewRotor(RotorII, A, A),
				NewRotor(RotorI, A, A),
			},
			ref:   NewReflector(ReflectorB),
			chars: []Char{A, A, A, A, A},
			res:   []Char{B, D, Z, G, O},
		},
		{
			name: "real data - with start position",
			rot: []Rotor{
				NewRotor(RotorIII, A, B),
				NewRotor(RotorII, A, B),
				NewRotor(RotorI, A, B),
			},
			ref:   NewReflector(ReflectorB),
			chars: []Char{A, A, A, A, A},
			res:   []Char{P, G, Q, P, W},
		},
		{
			name: "real data - with ring settings",
			rot: []Rotor{
				NewRotor(RotorIII, B, A),
				NewRotor(RotorII, B, A),
				NewRotor(RotorI, B, A),
			},
			ref:   NewReflector(ReflectorB),
			chars: []Char{A, A, A, A, A},
			res:   []Char{E, W, T, Y, X},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpindle(tt.rot, tt.ref)

			res := make([]Char, 0, len(tt.chars))
			for _, c := range tt.chars {
				cc := s.Handle(c)
				res = append(res, cc)
			}
			require.Equal(t, tt.res, res)
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name      string
		rotors    []Rotor
		steps     int
		positions []Char
	}{
		{
			name:      "one rotor",
			rotors:    []Rotor{NewRotor(RotorI, A, A)},
			steps:     3,
			positions: []Char{D},
		},
		{
			name: "three rotors - turnover on first rotor",
			rotors: []Rotor{
				NewRotor(RotorI, A, A),
				NewRotor(RotorII, A, A),
				NewRotor(RotorIII, A, A),
			},
			steps:     3,
			positions: []Char{D, A, A},
		},
		{
			name: "three rotors - turnover on first and second rotor",
			rotors: []Rotor{
				NewRotor(RotorI, A, P),
				NewRotor(RotorII, A, A),
				NewRotor(RotorIII, A, A),
			},
			steps:     3,
			positions: []Char{S, B, A},
		},
		{
			name: "three rotors - turnover on all rotors",
			rotors: []Rotor{
				NewRotor(RotorI, A, Q),
				NewRotor(RotorII, A, E),
				NewRotor(RotorIII, A, A),
			},
			steps:     1,
			positions: []Char{R, F, B},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.steps; i++ {
				rotate(tt.rotors)
			}

			for i := 0; i < len(tt.rotors); i++ {
				require.Equal(t, tt.positions[i], tt.rotors[i].Position())
			}
		})
	}
}
