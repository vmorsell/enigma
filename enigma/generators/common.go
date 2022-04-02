package generators

import (
	"fmt"

	"github.com/vmorsell/enigma/enigma"
)

func KeyMap(mapping string) (map[enigma.Key]enigma.Key, error) {
	alphas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(alphas) != len(mapping) {
		return nil, fmt.Errorf("not enough keys")
	}

	from, err := enigma.StringToKeys(alphas)
	if err != nil {
		return nil, fmt.Errorf("alphas string to keys: %w", err)
	}

	to, err := enigma.StringToKeys(mapping)
	if err != nil {
		return nil, fmt.Errorf("mapping string to keys: %w", err)
	}

	out := make(map[enigma.Key]enigma.Key, len(from))
	for i := 0; i < len(from); i++ {
		out[from[i]] = to[i]
	}
	return out, nil
}
