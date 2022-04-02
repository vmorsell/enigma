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

func (k Key) Shift(offset int) Key {
	new := (k + Key(offset))
	for new < 0 {
		new += Key(len(values))
	}
	return new % Key(len(values))
}

func (k Key) String() string {
	return values[k]
}

var errUnknownKey = func(s string) error { return fmt.Errorf("unknown key %s", s) }

func stringToKeys(s string) ([]Key, error) {
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
