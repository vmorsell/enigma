package generators

import (
	"fmt"

	"github.com/vmorsell/enigma/enigma"
)

// CharMap creates a mapping from an ordered A-Z string to the provided mapping.
func CharMap(mapping string) (map[enigma.Char]enigma.Char, error) {
	alphas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(alphas) != len(mapping) {
		return nil, fmt.Errorf("not enough chars")
	}

	from, err := enigma.StringToChars(alphas)
	if err != nil {
		return nil, fmt.Errorf("alphas string to chars: %w", err)
	}

	to, err := enigma.StringToChars(mapping)
	if err != nil {
		return nil, fmt.Errorf("mapping string to chars: %w", err)
	}

	out := make(map[enigma.Char]enigma.Char, len(from))
	for i := 0; i < len(from); i++ {
		out[from[i]] = to[i]
	}
	return out, nil
}

// BackwarddCharMap creates a reverse mapping from each letter in the provided
// mapping string, to the corresponding letter in an ordered A-Z string.
func BackwardCharMap(mapping string) (map[enigma.Char]enigma.Char, error) {
	fw, err := CharMap(mapping)
	if err != nil {
		return nil, fmt.Errorf("char map: %w", err)
	}

	return reverseMap(fw), nil
}

// reverseMap reverses the provided Char map k->v to v->k.
func reverseMap(m map[enigma.Char]enigma.Char) map[enigma.Char]enigma.Char {
	res := make(map[enigma.Char]enigma.Char, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}
