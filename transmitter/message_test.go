package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	s := "Hello, world!"
	want := "HELLOWORLD"

	got, err := format(s)
	require.Nil(t, err)
	require.Equal(t, want, got)
}
