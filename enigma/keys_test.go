package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	k := C
	want := 2

	got := k.Int()
	require.Equal(t, want, got)
}

func TestShift(t *testing.T) {
	tests := []struct {
		name   string
		k      Key
		offset int
		res    Key
	}{
		{
			name:   "positive offset",
			k:      A,
			offset: 1,
			res:    B,
		},
		{
			name:   "negative offset",
			k:      B,
			offset: -1,
			res:    A,
		},
		{
			name:   "overflow",
			k:      A,
			offset: 27,
			res:    B,
		},
		{
			name:   "underflow",
			k:      A,
			offset: -1,
			res:    Z,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.k.Shift(tt.offset)
			require.Equal(t, tt.res, res)
		})
	}
}
