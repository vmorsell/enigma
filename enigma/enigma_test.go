package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		k    Key
		res  Key
		err  error
	}{
		{
			name: "ok",
			k:    X,
			res:  X,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			e.Plugboard.Patch(X, Y)
			res, err := e.Encrypt(tt.k)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.res, res)
		})
	}
}
