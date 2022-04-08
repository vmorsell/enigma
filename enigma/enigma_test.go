package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptMessage(t *testing.T) {
	tests := []struct {
		name string
		msg  []Char
		dk   DailyKey
		mk   MessageKey
		out  EncryptedMessage
	}{
		{
			name: "ok",
			msg:  []Char{H, E, L, L, O},
			dk: DailyKey{
				RotorTypes:    []RotorType{RotorI, RotorII, RotorIII},
				ReflectorType: ReflectorA,
				Rings:         []Char{A, A, A},
				Positions:     []Char{A, A, A},
			},
			mk: MessageKey{
				Positions: []Char{A, A, A},
			},
			out: EncryptedMessage{
				EncryptedMessageKey: MessageKey{
					Positions: []Char{S, Q, C, X, F, K},
				},
				Payload: []Char{L, Z, F, B, D},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnigma(tt.dk)
			out := e.EncryptMessage(tt.msg, tt.dk, tt.mk)
			require.Equal(t, tt.out, out)
		})
	}
}

func TestDecryptMessage(t *testing.T) {
	tests := []struct {
		name string
		msg  EncryptedMessage
		dk   DailyKey
		out  []Char
	}{
		{
			name: "ok",
			msg: EncryptedMessage{
				EncryptedMessageKey: MessageKey{
					Positions: []Char{S, Q, C, F, X, K},
				},
				Payload: []Char{L, Z, F, B, D},
			},
			dk: DailyKey{
				RotorTypes:    []RotorType{RotorI, RotorII, RotorIII},
				ReflectorType: ReflectorA,
				Rings:         []Char{A, A, A},
				Positions:     []Char{A, A, A},
			},
			out: []Char{H, E, L, L, O},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnigma(tt.dk)
			out := e.DecryptMessage(tt.msg, tt.dk)
			require.Equal(t, tt.out, out)
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := DailyKey{
		RotorTypes:    []RotorType{RotorIII, RotorI, RotorII},
		ReflectorType: ReflectorA,
		Rings:         []Char{V, M, X},
		Positions:     []Char{L, B, A},
		Plugs: []PlugboardMapping{
			{A, M},
			{F, I},
			{N, V},
			{P, S},
			{T, U},
			{W, Z},
		},
	}
	e := NewEnigma(key)

	in := []Char{S, E, C, R, E, T}
	want := []Char{L, C, G, O, D, U}

	got := e.EncryptDecrypt(in)
	require.Equal(t, want, got)
}

func TestEncryptDecryptChar(t *testing.T) {
	key := DailyKey{
		RotorTypes:    []RotorType{RotorI, RotorII, RotorIII},
		ReflectorType: ReflectorA,
		Rings:         []Char{A, B, C},
		Positions:     []Char{C, B, A},
		Plugs: []PlugboardMapping{
			{A, X},
		},
	}
	e := NewEnigma(key)

	in := A
	want := J

	res := e.(*enigma).encryptDecryptChar(in)
	require.Equal(t, want, res)
}
