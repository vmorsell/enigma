package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	k := C
	want := 2

	got := k.Int()
	require.Equal(t, want, got)
}

func TestStepKey(t *testing.T) {
	k := C
	want := D

	got := k.Step()
	require.Equal(t, want, got)
}

func TestShift(t *testing.T) {
	tests := []struct {
		name   string
		k      Key
		offset Key
		res    Key
	}{
		{
			name:   "positive offset",
			k:      A,
			offset: C,
			res:    C,
		},
		{
			name:   "negative offset",
			k:      B,
			offset: -B,
			res:    A,
		},
		{
			name:   "overflow",
			k:      Z,
			offset: C,
			res:    B,
		},
		{
			name:   "underflow",
			k:      A,
			offset: -B,
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

func TestString(t *testing.T) {
	alphaOffset := 65

	for i := 0; i < len(values); i++ {
		want := fmt.Sprintf("%c", i+alphaOffset)
		got := Key(i).String()
		require.Equal(t, want, got)
	}
}

func TestStringToKeys(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    []Key
		err  error
	}{
		{
			name: "not ok - key not found",
			s:    "Ö",
			err:  errUnknownKey("Ö"),
		},
		{
			name: "ok",
			s:    "A",
			k:    []Key{A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, err := StringToKeys(tt.s)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.k, keys)
		})
	}
}
