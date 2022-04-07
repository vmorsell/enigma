package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageKeyEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name string
		dk   DailyKey
		in   MessageKey
		out  MessageKey
	}{
		{
			name: "ok",
			dk: DailyKey{
				RotorTypes:    []RotorType{RotorI, RotorII, RotorIII},
				ReflectorType: ReflectorA,
				Rings:         []Char{A, A, A},
				Positions:     []Char{A, A, A},
			},
			in: MessageKey{
				Positions: []Char{A, A, A},
			},
			out: MessageKey{
				Positions: []Char{S, Q, C},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnigma(tt.dk)

			t.Run("encrypt", func(t *testing.T) {
				res := tt.in.EncryptDecrypt(e)
				require.Equal(t, tt.out, res)
			})

			t.Run("decrypt", func(t *testing.T) {
				e.SetDailyKey(tt.dk)
				res := tt.out.EncryptDecrypt(e)
				require.Equal(t, tt.in, res)
			})
		})
	}
}
