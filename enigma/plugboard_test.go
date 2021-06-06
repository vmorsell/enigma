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
			a string
			b string
		}
		v   string
		res string
	}{
		{
			name: "without patches",
			v:    "x",
			res:  "x",
		},
		{
			name: "with patch",
			patch: &struct {
				a string
				b string
			}{"x", "y"},
			v:   "x",
			res: "y",
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
		patches map[string]string
		a       string
		b       string
		err     error
	}{
		{
			name:    "a already patched",
			patches: map[string]string{"x": "a"},
			a:       "x",
			b:       "y",
			err:     fmt.Errorf("x already patched to a"),
		},
		{
			name:    "b already patched",
			patches: map[string]string{"y": "a"},
			a:       "x",
			b:       "y",
			err:     fmt.Errorf("y already patched to a"),
		},
		{
			name:    "ok",
			patches: make(map[string]string),
			a:       "x",
			b:       "y",
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
