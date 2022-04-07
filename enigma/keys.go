package enigma

import "math/rand"

// DailyKey is the initialization vector for the Enigma machine. It's used
// as base settings for all encryption and decryption.
type DailyKey struct {
	Rotors          []RotorType
	Reflector       ReflectorType
	RingSettings    []Char
	RotorPositions  []Char
	PlugConnections []PlugboardMapping
}

// NewDailyKey returs a daily key.
func NewDailyKey(rotors []RotorType, reflector ReflectorType, ringSettings []Char, rotorPositions []Char, plugConnections []PlugboardMapping) DailyKey {
	return DailyKey{
		Rotors:          rotors,
		Reflector:       reflector,
		RingSettings:    ringSettings,
		RotorPositions:  rotorPositions,
		PlugConnections: plugConnections,
	}
}

// Random daily key returns a randomized daily key.
func RandomDailyKey() DailyKey {
	rot := randomRotors(3)
	ref := randomReflector()
	ring := randomChars(3)
	pos := randomChars(3)
	plug := randomPlugConnections(10)

	return NewDailyKey(rot, ref, ring, pos, plug)
}

// randomRotors returns n rotor types in a random order.
func randomRotors(n int) []RotorType {
	all := AllRotorTypes[:]
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})
	return all[:n]
}

// randomReflector returns a random reflector type.
func randomReflector() ReflectorType {
	all := AllReflectorTypes[:]
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})
	return all[0]
}

// randomChars returns n randomized, non-unique chars.
func randomChars(n int) []Char {
	max := len(charSet)
	chars := make([]Char, 0, n)
	for len(chars) < n {
		r := Char(rand.Intn(max))
		chars = append(chars, r)
	}
	return chars
}

// randomPlugConnections returns n random plug connections. All connections
// are to different chars.
func randomPlugConnections(n int) []PlugboardMapping {
	all := make([]int, 0, len(charSet))
	for i := 0; i < len(charSet); i++ {
		all = append(all, i)
	}
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	mappings := make([]PlugboardMapping, 0, n)
	for i := 0; i < n*2; i += 2 {
		m := PlugboardMapping{
			From: Char(all[i]),
			To:   Char(all[i+1]),
		}
		mappings = append(mappings, m)
	}
	return mappings
}
