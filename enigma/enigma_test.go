package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		v    string
		res  string
		err  error
	}{
		{
			name: "multiple characters",
			v:    "ab",
			err:  fmt.Errorf("must encrypt one character a time"),
		},
		{
			name: "ok",
			v:    "x",
			res:  "X",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			e.Plugboard.Patch("x", "y")
			res, err := e.Encrypt(tt.v)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.res, res)
		})
	}
}
