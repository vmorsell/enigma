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
			rot: []Rotor{NewRotor(rotorConfig{
				mapping: map[Key]Key{
					A: B,
					Y: X,
				},
				notch: Z,
			}, A)},
			ref: NewReflector(reflectorConfig{
				mapping: map[Key]Key{
					B: X,
				},
			}),
			keys: []Key{A},
			res:  []Key{Y},
		},
		{
			name: "one rotor - multiple keys",
			rot: []Rotor{NewRotor(rotorConfig{
				mapping: map[Key]Key{
					A: R,
					J: B,
					Z: O,
					L: J,
				},
				notch: Z,
			}, A)},
			ref: NewReflector(reflectorConfig{
				mapping: map[Key]Key{
					R: B,
					P: K,
				},
			}),
			keys: []Key{A, A},
			res:  []Key{J, M},
		},
		{
			name: "two rotors - multiple keys",
			rot: []Rotor{
				NewRotor(rotorConfig{
					mapping: map[Key]Key{
						A: R,
						V: U,

						Z: G,
						S: O,
					},
					notch: Z,
				}, A),
				NewRotor(rotorConfig{
					mapping: map[Key]Key{
						R: T,
						U: D,

						H: N,
						P: O,
					},
					notch: Z,
				}, A),
			},
			ref: NewReflector(reflectorConfig{
				mapping: map[Key]Key{
					T: D,
					N: O,
				},
			}),
			keys: []Key{A, A},
			res:  []Key{V, T},
		},
		{
			name: "two rotors - multiple keys - both rotors rotates after first key",
			rot: []Rotor{
				NewRotor(rotorConfig{
					mapping: map[Key]Key{
						A: T,
						E: R,

						Z: L,
						K: F,
					},
					notch: A,
				}, A),
				NewRotor(rotorConfig{
					mapping: map[Key]Key{
						T: Q,
						R: G,

						L: U,
						F: S,
					},
					notch: Z,
				}, A),
			},
			ref: NewReflector(reflectorConfig{
				mapping: map[Key]Key{
					Q: G,
					V: T,
				},
			}),
			keys: []Key{A, A},
			res:  []Key{E, L},
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
