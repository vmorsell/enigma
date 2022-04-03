package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpindleHandle(t *testing.T) {
	tests := []struct {
		name string
		rot  []Rotor
		ref  Reflector
		keys []Key
		res  []Key
	}{
		{
			name: "one rotor - one key",
			rot:  []Rotor{NewRotor(RotorI, A, A)},
			ref:  NewReflector(ReflectorA),
			keys: []Key{A},
			res:  []Key{X},
		},
		{
			name: "one rotor - multiple keys",
			rot:  []Rotor{NewRotor(RotorI, A, A)},
			ref:  NewReflector(ReflectorA),
			keys: []Key{A, A},
			res:  []Key{X, M},
		},
		{
			name: "two rotors - multiple keys",
			rot: []Rotor{
				NewRotor(RotorI, A, A),
				NewRotor(RotorII, A, A),
			},
			ref:  NewReflector(ReflectorA),
			keys: []Key{A, A},
			res:  []Key{X, M},
		},
		{
			name: "three rotors",
			rot: []Rotor{
				NewRotor(RotorIII, A, A),
				NewRotor(RotorII, A, A),
				NewRotor(RotorI, A, A),
			},
			ref:  NewReflector(ReflectorB),
			keys: []Key{A, A, A, A, A},
			res:  []Key{B, D, Z, G, O},
		},
		{
			name: "real data - with start position",
			rot: []Rotor{
				NewRotor(RotorIII, A, B),
				NewRotor(RotorII, A, B),
				NewRotor(RotorI, A, B),
			},
			ref:  NewReflector(ReflectorB),
			keys: []Key{A, A, A, A, A},
			res:  []Key{P, G, Q, P, W},
		},
		{
			name: "real data - with ring settings",
			rot: []Rotor{
				NewRotor(RotorIII, B, A),
				NewRotor(RotorII, B, A),
				NewRotor(RotorI, B, A),
			},
			ref:  NewReflector(ReflectorB),
			keys: []Key{A, A, A, A, A},
			res:  []Key{E, W, T, Y, X},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpindle(tt.rot, tt.ref)

			res := make([]Key, 0, len(tt.keys))
			for _, k := range tt.keys {
				kk := s.Handle(k)
				res = append(res, kk)
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
		positions []Key
	}{
		{
			name:      "one rotor",
			rotors:    []Rotor{NewRotor(RotorI, A, A)},
			steps:     3,
			positions: []Key{D},
		},
		{
			name: "three rotors - turnover on first rotor",
			rotors: []Rotor{
				NewRotor(RotorI, A, A),
				NewRotor(RotorII, A, A),
				NewRotor(RotorIII, A, A),
			},
			steps:     3,
			positions: []Key{D, A, A},
		},
		{
			name: "three rotors - turnover on first and second rotor",
			rotors: []Rotor{
				NewRotor(RotorI, A, P),
				NewRotor(RotorII, A, A),
				NewRotor(RotorIII, A, A),
			},
			steps:     3,
			positions: []Key{S, B, A},
		},
		{
			name: "three rotors - turnover on all rotors",
			rotors: []Rotor{
				NewRotor(RotorI, A, Q),
				NewRotor(RotorII, A, E),
				NewRotor(RotorIII, A, A),
			},
			steps:     1,
			positions: []Key{R, F, B},
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
