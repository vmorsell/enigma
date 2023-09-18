package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReflect(t *testing.T) {
	r := NewReflector(ReflectorType{
		mapping: map[Char]Char{
			A: B,
		},
	})

	tests := []struct {
		name string
		c    Char
		res  Char
	}{
		{
			name: "found in map",
			c:    A,
			res:  B,
		},
		{
			name: "not found",
			c:    X,
			res:  X,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := r.Handle(tt.c)
			require.Equal(t, tt.res, res)
		})
	}
}
