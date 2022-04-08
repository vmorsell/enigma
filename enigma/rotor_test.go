package enigma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRotor(t *testing.T) {
	tests := []struct {
		name         string
		typ          RotorType
		ring         Char
		position     Char
		wantPosition Char
		wantNotch    Char
	}{
		{
			name:         "no ring or position",
			typ:          RotorI,
			wantPosition: A,
			wantNotch:    RotorI.notch,
		},
		{
			name:         "with initial position",
			typ:          RotorI,
			position:     D,
			wantPosition: D,
			wantNotch:    RotorI.notch,
		},
		{
			name:         "with changed ring",
			typ:          RotorI,
			position:     G,
			ring:         B,
			wantPosition: F,
			wantNotch:    RotorI.notch - B,
		},
		{
			name:         "with initial position and changed ring",
			typ:          RotorI,
			position:     D,
			ring:         D,
			wantPosition: A,
			wantNotch:    RotorI.notch - D,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRotor(tt.typ, tt.ring, tt.position)
			require.Equal(t, tt.wantPosition, r.Position(), "position")
			require.Equal(t, tt.wantNotch, r.Notch(), "notch")
		})
	}
}

func TestForwardAndBackward(t *testing.T) {
	tests := []struct {
		name string
		rot  Rotor
		c    Char
		out  Char
	}{
		{
			name: "ok",
			rot:  NewRotor(RotorI, A, A),
			c:    A,
			out:  E,
		},
		{
			name: "with ring setting",
			rot:  NewRotor(RotorI, B, A),
			c:    A,
			out:  K,
		},
		{
			name: "with shifted position",
			rot:  NewRotor(RotorI, A, B),
			c:    A,
			out:  J,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("forward", func(t *testing.T) {
				out := tt.rot.Forward(tt.c)
				require.Equal(t, tt.out, out)
			})

			t.Run("backward", func(t *testing.T) {
				c := tt.rot.Backward(tt.out)
				require.Equal(t, tt.c, c)
			})
		})
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		name string
		init Char
		want Char
	}{
		{
			name: "ok - from A",
			init: A,
			want: B,
		},
		{
			name: "ok - from non-A",
			init: G,
			want: H,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRotor(RotorI, A, tt.init)
			r.Step()
			require.Equal(t, tt.want, r.Position())
		})
	}
}

func TestSetRing(t *testing.T) {
	ring := B
	pos := A
	notch := K

	wantPos := Z
	wantNotch := J

	r := rotor{
		typ: RotorType{
			notch: notch,
		},
		position: pos,
	}
	r.SetRing(ring)
	require.Equal(t, wantPos, r.Position(), "position")
	require.Equal(t, wantNotch, r.Notch(), "notch")
}

func TestSetPosition(t *testing.T) {
	tests := []struct {
		name string
		init Char
		want Char
	}{
		{
			name: "ok - from A",
			init: A,
			want: C,
		},
		{
			name: "ok - from non-A",
			init: G,
			want: H,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRotor(RotorI, A, tt.init)
			r.SetPosition(tt.want)
			require.Equal(t, tt.want, r.Position())
		})
	}
}
