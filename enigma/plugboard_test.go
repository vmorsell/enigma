package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		name  string
		patch *struct {
			a Key
			b Key
		}
		v   Key
		res Key
	}{
		{
			name: "without patches",
			v:    X,
			res:  X,
		},
		{
			name: "with patch",
			patch: &struct {
				a Key
				b Key
			}{X, Y},
			v:   X,
			res: Y,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPlugboard()
			if tt.patch != nil {
				p.Patch(tt.patch.a, tt.patch.b)
			}
			res := p.Handle(tt.v)
			require.Equal(t, tt.res, res)
		})
	}
}

func TestPatch(t *testing.T) {
	tests := []struct {
		name    string
		patches map[Key]Key
		a       Key
		b       Key
		err     error
	}{
		{
			name:    "a already patched",
			patches: map[Key]Key{X: A},
			a:       X,
			b:       Y,
			err:     fmt.Errorf("X already patched to A"),
		},
		{
			name:    "B already patched",
			patches: map[Key]Key{Y: A},
			a:       X,
			b:       Y,
			err:     fmt.Errorf("Y already patched to A"),
		},
		{
			name:    "ok",
			patches: make(map[Key]Key),
			a:       X,
			b:       Y,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugboard{
				Patches: tt.patches,
			}
			err := p.Patch(tt.a, tt.b)
			require.Equal(t, tt.err, err)

			if err == nil {
				require.Equal(t, tt.b, p.Patches[tt.a])
				require.Equal(t, tt.a, p.Patches[tt.b])
			}
		})
	}
}
