package enigma

type Enigma struct {
	plugboard Plugboard
	spindle   Spindle
}

// New returns an Enigma instance.
func New(plugboard Plugboard, spindle Spindle) Enigma {
	return Enigma{
		plugboard: plugboard,
		spindle:   spindle,
	}
}

// Encrypt encrypts a slice of keys.
func (e Enigma) Encrypt(keys []Key) []Key {
	res := make([]Key, 0, len(keys))
	for _, k := range keys {
		kk := e.encryptKey(k)
		res = append(res, kk)
	}
	return res
}

// encryptKey encrypts a single key.
func (e Enigma) encryptKey(k Key) Key {
	k = e.plugboard.Handle(k)
	k = e.spindle.Handle(k)
	k = e.plugboard.Handle(k)
	return k
}
