package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	e := New()
	e.Plugboard.Patch("x", "y")
	res := e.Encrypt("x")
	require.Equal(t, "x", res)
}
