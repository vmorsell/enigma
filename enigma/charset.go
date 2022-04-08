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

// Int returns the char as an int.
func (c Char) Int() int {
	return int(c)
}

// Step increments the char one step.
func (c Char) Step() Char {
	return c.Shift(1)
}

// Shift shifts the char the provided amount of steps.
func (c Char) Shift(char Char) Char {
	c += char
	for c < 0 {
		c += Char(len(charSet))
	}
	return c % Char(len(charSet))
}

// String returns the char as at string representation.
func (c Char) String() string {
	return charSet[c]
}

var errNotInCharset = func(s string) error { return fmt.Errorf("%s is not in charset", s) }

// StringToChars converts a string to a slice with chars.
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

// CharsToString converts a char slice to a string.
func CharsToString(chars []Char) string {
	var out string
	for _, c := range chars {
		out += c.String()
	}
	return out
}
