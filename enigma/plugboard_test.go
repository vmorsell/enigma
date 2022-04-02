package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlugboardForwardAndBackward(t *testing.T) {
	forwardMap := map[Key]Key{
		A: B,
	}
	pb := NewPlugboard(forwardMap)

	t.Run("forward", func(t *testing.T) {

		tests := []struct {
			name string
			k    Key
			res  Key
		}{
			{
				name: "key mapped",
				k:    A,
				res:  B,
			},
			{
				name: "key not mapped",
				k:    X,
				res:  X,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := pb.Forward(tt.k)
				require.Equal(t, tt.res, res)
			})
		}
	})

	t.Run("backward", func(t *testing.T) {

		tests := []struct {
			name string
			k    Key
			res  Key
		}{
			{
				name: "key mapped",
				k:    B,
				res:  A,
			},
			{
				name: "key not mapped",
				k:    X,
				res:  X,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := pb.Backward(tt.k)
				require.Equal(t, tt.res, res)
			})
		}
	})
}
