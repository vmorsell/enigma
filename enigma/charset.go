package enigma

import "fmt"

type Char int

const (
	A Char = iota
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

var charSet = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func (c Char) Int() int {
	return int(c)
}

func (c Char) Step() Char {
	return c.Shift(1)
}

func (c Char) Shift(char Char) Char {
	c += char
	for c < 0 {
		c += Char(len(charSet))
	}
	return c % Char(len(charSet))
}

func (c Char) String() string {
	return charSet[c]
}

var errNotInCharset = func(s string) error { return fmt.Errorf("%s is not in charset", s) }

func StringToChars(s string) ([]Char, error) {
	out := make([]Char, 0, len(s))
nextRune:
	for _, r := range s {
		c := fmt.Sprintf("%c", r)
		for i, v := range charSet {
			if v == c {
				out = append(out, Char(i))
				continue nextRune
			}
		}
		return nil, errNotInCharset(s)
	}
	return out, nil
}

func CharsToString(chars []Char) string {
	var out string
	for _, c := range chars {
		out += c.String()
	}
	return out
}
