package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	pb := NewPlugboard(map[Key]Key{
		A: B,
	})

	tests := []struct {
		name string
		k    Key
		want Key
	}{
		{
			name: "not mapped",
			k:    X,
			want: X,
		},
		{
			name: "mapped as key",
			k:    A,
			want: B,
		},
		{
			name: "mapped as value",
			k:    B,
			want: A,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pb.Handle(tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}
