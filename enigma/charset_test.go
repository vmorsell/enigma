package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	c := C
	want := 2

	got := c.Int()
	require.Equal(t, want, got)
}

func TestStepChar(t *testing.T) {
	c := C
	want := D

	got := c.Step()
	require.Equal(t, want, got)
}

func TestShift(t *testing.T) {
	tests := []struct {
		name   string
		c      Char
		offset Char
		res    Char
	}{
		{
			name:   "positive offset",
			c:      A,
			offset: C,
			res:    C,
		},
		{
			name:   "negative offset",
			c:      B,
			offset: -B,
			res:    A,
		},
		{
			name:   "overflow",
			c:      Z,
			offset: C,
			res:    B,
		},
		{
			name:   "underflow",
			c:      A,
			offset: -B,
			res:    Z,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.c.Shift(tt.offset)
			require.Equal(t, tt.res, res)
		})
	}
}

func TestString(t *testing.T) {
	alphaOffset := 65

	for i := 0; i < len(charSet); i++ {
		want := fmt.Sprintf("%c", i+alphaOffset)
		got := Char(i).String()
		require.Equal(t, want, got)
	}
}

func TestStringToChars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		c    []Char
		err  error
	}{
		{
			name: "not ok - not in charset",
			s:    "Ö",
			err:  errNotInCharset("Ö"),
		},
		{
			name: "ok",
			s:    "A",
			c:    []Char{A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chars, err := StringToChars(tt.s)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.c, chars)
		})
	}
}

func TestCharsToString(t *testing.T) {
	in := []Char{A, B, C}
	want := "ABC"
	got := CharsToString(in)
	require.Equal(t, want, got)
}
