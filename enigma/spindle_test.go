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
			rot:  []Rotor{NewRotor(map[Key]Key{A: B, Y: X}, A, Z)},
			ref:  NewReflector(map[Key]Key{B: X}),
			keys: []Key{A},
			res:  []Key{Y},
		},
		{
			name: "one rotor - multiple keys",
			rot: []Rotor{NewRotor(map[Key]Key{
				A: R,
				J: B,
				Z: O,
				L: J,
			}, A, Z)},
			ref: NewReflector(map[Key]Key{
				R: B,
				P: K,
			}),
			keys: []Key{A, A},
			res:  []Key{J, M},
		},
		{
			name: "two rotors - multiple keys",
			rot: []Rotor{
				NewRotor(map[Key]Key{
					A: R,
					V: U,

					Z: G,
					S: O,
				}, A, Z),
				NewRotor(map[Key]Key{
					R: T,
					U: D,

					H: N,
					P: O,
				}, A, Z),
			},
			ref: NewReflector(map[Key]Key{
				T: D,
				N: O,
			}),
			keys: []Key{A, A},
			res:  []Key{V, T},
		},
		{
			name: "two rotors - multiple keys - both rotors rotates after first key",
			rot: []Rotor{
				NewRotor(map[Key]Key{
					A: T,
					E: R,

					Z: L,
					K: F,
				}, A, A),
				NewRotor(map[Key]Key{
					T: Q,
					R: G,

					L: U,
					F: S,
				}, A, Z),
			},
			ref: NewReflector(map[Key]Key{
				Q: G,
				V: T,
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
