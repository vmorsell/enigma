package generators

import (
	"fmt"

	"github.com/vmorsell/enigma/enigma"
)

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
