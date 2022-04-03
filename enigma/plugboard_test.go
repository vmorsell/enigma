package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlugboardMaps(t *testing.T) {
	mapping := []PlugboardMapping{
		{A, S},
		{Z, T},
	}
	wantFw := map[Char]Char{
		A: S,
		Z: T,
	}
	wantBw := map[Char]Char{
		S: A,
		T: Z,
	}

	fw, bw := plugboardMaps(mapping)
	require.EqualValues(t, wantFw, fw, "forward")
	require.EqualValues(t, wantBw, bw, "backward")
}

func TestHandle(t *testing.T) {
	pb := NewPlugboard(PlugboardSettings{
		Mappings: []PlugboardMapping{
			{A, B},
		},
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
			name: "mapped forward",
			c:    A,
			want: B,
		},
		{
			name: "mapped backward",
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
