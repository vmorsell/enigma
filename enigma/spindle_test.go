package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpindleHandle(t *testing.T) {
	tests := []struct {
		name      string
		rotors    []RotorType
		reflector ReflectorType
		rings     []Char
		positions []Char
		in        []Char
		out       []Char
	}{
		{
			name:      "one rotor - one char",
			rotors:    []RotorType{RotorI},
			reflector: ReflectorA,
			rings:     []Char{A},
			positions: []Char{A},
			in:        []Char{A},
			out:       []Char{X},
		},
		{
			name:      "one rotor - multiple chars",
			rotors:    []RotorType{RotorI},
			reflector: ReflectorA,
			rings:     []Char{A},
			positions: []Char{A},
			in:        []Char{A, A},
			out:       []Char{X, M},
		},
		{
			name:      "two rotors - multiple chars",
			rotors:    []RotorType{RotorI, RotorII},
			reflector: ReflectorA,
			rings:     []Char{A, A},
			positions: []Char{A, A},
			in:        []Char{A, A},
			out:       []Char{X, M},
		},
		{
			name:      "three rotors",
			rotors:    []RotorType{RotorIII, RotorII, RotorI},
			reflector: ReflectorB,
			rings:     []Char{A, A, A},
			positions: []Char{A, A, A},
			in:        []Char{A, A, A, A, A},
			out:       []Char{B, D, Z, G, O},
		},
		{
			name:      "three rotors - with start position",
			rotors:    []RotorType{RotorIII, RotorII, RotorI},
			reflector: ReflectorB,
			rings:     []Char{A, A, A},
			positions: []Char{B, B, B},
			in:        []Char{A, A, A, A, A},
			out:       []Char{P, G, Q, P, W},
		},
		{
			name:      "three rotors - with ring settings",
			rotors:    []RotorType{RotorIII, RotorII, RotorI},
			reflector: ReflectorB,
			rings:     []Char{B, B, B},
			positions: []Char{A, A, A},
			in:        []Char{A, A, A, A, A},
			out:       []Char{E, W, T, Y, X},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpindle(tt.rotors, tt.reflector, tt.rings, tt.positions)

			out := make([]Char, 0, len(tt.in))
			for _, c := range tt.in {
				cc := s.Handle(c)
				out = append(out, cc)
			}
			require.Equal(t, tt.out, out)
		})
	}
}

func TestSetPositions(t *testing.T) {
	tests := []struct {
		name       string
		rotorTypes []RotorType
		positions  []Char
		want       []Char
	}{
		{
			name:       "ok - from A",
			rotorTypes: []RotorType{RotorI, RotorI, RotorI},
			positions:  []Char{A, A, A},
			want:       []Char{A, B, C},
		},
		{
			name:       "ok - not from A",
			rotorTypes: []RotorType{RotorI, RotorI, RotorI},
			positions:  []Char{B, C, D},
			want:       []Char{E, A, O},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpindle(tt.rotorTypes, ReflectorA, tt.positions, tt.positions)
			s.SetPositions(tt.want)

			var got []Char
			for _, r := range s.(*spindle).rotors {
				got = append(got, r.Position())
			}
			require.Equal(t, tt.want, got)
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
