package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	pb := NewPlugboard(map[Char]Char{
		A: B,
	})

	tests := []struct {
		name string
		c    Char
		want Char
	}{
		{
			name: "not mapped",
			c:    X,
			want: X,
		},
		{
			name: "mapped as char",
			c:    A,
			want: B,
		},
		{
			name: "mapped as value",
			c:    B,
			want: A,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pb.Handle(tt.c)
			require.Equal(t, tt.want, got)
		})
	}
}
