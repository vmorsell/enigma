package enigma

import "fmt"

type Key int

const (
	A Key = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

var values = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func (k Key) Int() int {
	return int(k)
}

func (k Key) Step() Key {
	return k.Shift(1)
}

func (k Key) Shift(key Key) Key {
	k += key
	for k < 0 {
		k += Key(len(values))
	}
	return k % Key(len(values))
}

func (k Key) String() string {
	return values[k]
}

var errUnknownKey = func(s string) error { return fmt.Errorf("unknown key %s", s) }

func StringToKeys(s string) ([]Key, error) {
	out := make([]Key, 0, len(s))
nextRune:
	for _, r := range s {
		c := fmt.Sprintf("%c", r)
		for i, v := range values {
			if v == c {
				out = append(out, Key(i))
				continue nextRune
			}
		}
		return nil, errUnknownKey(s)
	}
	return out, nil
}

func keysToString(keys []Key) string {
	var out string
	for _, k := range keys {
		out += k.String()
	}
	return out
}
