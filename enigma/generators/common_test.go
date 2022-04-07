package generators

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vmorsell/enigma/enigma"
)

func TestReverseMap(t *testing.T) {
	in := map[enigma.Char]enigma.Char{
		0: 1,
	}
	want := map[enigma.Char]enigma.Char{
		1: 0,
	}

	res := reverseMap(in)
	require.Equal(t, want, res)
}
