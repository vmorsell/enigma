package enigma

import "math/rand"

// DailyKey is the initialization vector for the Enigma machine. It's used
// as base settings for all encryption and decryption.
type DailyKey struct {
	RotorTypes    []RotorType
	ReflectorType ReflectorType
	Rings         []Char
	Positions     []Char
	Plugs         []PlugboardMapping
}

// NewDailyKey returs a daily key.
func NewDailyKey(rotorTypes []RotorType, reflectorType ReflectorType, rings []Char, positions []Char, plugs []PlugboardMapping) DailyKey {
	return DailyKey{
		RotorTypes:    rotorTypes,
		ReflectorType: reflectorType,
		Rings:         rings,
		Positions:     positions,
		Plugs:         plugs,
	}
}

// Random daily key returns a randomized daily key.
func RandomDailyKey() DailyKey {
	rot := randomRotors(3)
	ref := randomReflector()
	rng := randomChars(3)
	pos := randomChars(3)
	plg := randomPlugConnections(10)

	return NewDailyKey(rot, ref, rng, pos, plg)
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

// MessageKey is the complementary initialization vector used to encrypt and
// decrypt indivual messages.
type MessageKey struct {
	Positions []Char
}

// Encrypt encrypts a message key.
func (k MessageKey) Encrypt(e Enigma) MessageKey {
	// Encrypt the start positions two times.
	payload := append(k.Positions, k.Positions...)
	res := e.EncryptDecrypt(payload)

	return MessageKey{
		Positions: res,
	}
}

// Decrypt decrypts a message key.
func (k MessageKey) Decrypt(e Enigma) MessageKey {
	decrypted := e.EncryptDecrypt(k.Positions)

	// The key should have been decrypted two times.
	// Only return the first half.
	decrypted = decrypted[:3]

	return MessageKey{
		Positions: decrypted,
	}
}
