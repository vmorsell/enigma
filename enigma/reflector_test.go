package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReflect(t *testing.T) {
	r := NewReflector(reflectorConfig{
		mapping: map[Key]Key{
			A: B,
		},
	})

	tests := []struct {
		name string
		k    Key
		res  Key
	}{
		{
			name: "found in map",
			k:    A,
			res:  B,
		},
		{
			name: "not found",
			k:    X,
			res:  X,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := r.Reflect(tt.k)
			require.Equal(t, tt.res, res)
		})
	}
}
