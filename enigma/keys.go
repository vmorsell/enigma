package enigma

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

const numKeys = 26

func (k Key) Int() int {
	return int(k)
}

func (k Key) Shift(offset int) Key {
	new := (k + Key(offset))
	for new < 0 {
		new += numKeys
	}
	return new % numKeys
}
