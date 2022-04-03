package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRotor(t *testing.T) {
	tests := []struct {
		name         string
		config       rotorConfig
		ring         Key
		position     Key
		wantPosition Key
		wantNotch    Key
	}{
		{
			name:         "no ring or position",
			config:       RotorI,
			wantPosition: A,
			wantNotch:    RotorI.notch,
		},
		{
			name:         "with initial position",
			config:       RotorI,
			position:     D,
			wantPosition: D,
			wantNotch:    RotorI.notch,
		},
		{
			name:         "with changed ring",
			config:       RotorI,
			position:     G,
			ring:         B,
			wantPosition: F,
			wantNotch:    RotorI.notch - B,
		},
		{
			name:         "with initial position and changed ring",
			config:       RotorI,
			position:     D,
			ring:         D,
			wantPosition: A,
			wantNotch:    RotorI.notch - D,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRotor(tt.config, tt.ring, tt.position)
			require.Equal(t, tt.wantPosition, r.Position(), "position")
			require.Equal(t, tt.wantNotch, r.Notch(), "notch")
		})
	}
}
func TestReverseMap(t *testing.T) {
	in := map[Key]Key{A: B}
	want := map[Key]Key{B: A}

	res := reverseMap(in)
	require.Equal(t, want, res)
}

func TestForwardAndBackward(t *testing.T) {
	tests := []struct {
		name string
		rot  Rotor
		k    Key
		out  Key
	}{
		{
			name: "ok",
			rot:  NewRotor(RotorI, A, A),
			k:    A,
			out:  E,
		},
		{
			name: "with ring setting",
			rot:  NewRotor(RotorI, B, A),
			k:    A,
			out:  K,
		},
		{
			name: "with shifted position",
			rot:  NewRotor(RotorI, A, B),
			k:    A,
			out:  J,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("forward", func(t *testing.T) {
				out := tt.rot.Forward(tt.k)
				require.Equal(t, tt.out, out)
			})

			t.Run("backward", func(t *testing.T) {
				k := tt.rot.Backward(tt.out)
				require.Equal(t, tt.k, k)
			})
		})
	}
}

func TestStep(t *testing.T) {
	pos := A
	want := B
	r := NewRotor(rotorConfig{}, A, pos)

	r.Step()
	require.Equal(t, r.Position(), want)
}

func TestSetRing(t *testing.T) {
	ring := B
	pos := A
	notch := K

	wantPos := Z
	wantNotch := J

	r := rotor{
		position: pos,
		notch:    notch,
	}
	r.SetRing(ring)
	require.Equal(t, wantPos, r.Position(), "position")
	require.Equal(t, wantNotch, r.Notch(), "notch")
}
