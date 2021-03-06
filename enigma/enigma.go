package enigma

import "fmt"

const (
	payloadGroupSize = 5
)

// Enigma holds the interface for an Enigma instance.
type Enigma interface {
	EncryptMessage(payload []Char, dk DailyKey, mk MessageKey) EncryptedMessage
	DecryptMessage(msg EncryptedMessage, dk DailyKey) []Char
	EncryptDecrypt(chars []Char) []Char
	SetDailyKey(key DailyKey)
	SetMessageKey(key MessageKey)
}

// enigma holds the Enigma logic.
type enigma struct {
	plugboard Plugboard
	spindle   Spindle
}

// New returns an Enigma instance.
func NewEnigma(key DailyKey) Enigma {
	e := &enigma{}
	e.SetDailyKey(key)
	return e
}

// SetDailyKey applies a daily key for the Enigma instance.
func (e *enigma) SetDailyKey(key DailyKey) {
	pb := NewPlugboard(key.Plugs)
	spindle := NewSpindle(key.RotorTypes, key.ReflectorType, key.Rings, key.Positions)

	e.plugboard = pb
	e.spindle = spindle
}

// SetMessageKey applies a message key to the Enigma instance.
func (e *enigma) SetMessageKey(key MessageKey) {
	e.spindle.SetPositions(key.Positions)
}

// EncryptedMessage represents the resulting output of an Enigma encryption.
type EncryptedMessage struct {
	EncryptedMessageKey MessageKey
	Payload             []Char
}

// String returns a string representation of an encrypted message in the
// format as it should be transmitted.
func (msg EncryptedMessage) String() string {
	key := fmt.Sprintf("%s %s", CharsToString(msg.EncryptedMessageKey.Positions[:3]), CharsToString(msg.EncryptedMessageKey.Positions[3:]))

	payload := CharsToString(msg.Payload)
	for i := payloadGroupSize; i < len(payload); i += payloadGroupSize + 1 {
		payload = fmt.Sprintf("%s %s", payload[:i], payload[i:])
	}

	return fmt.Sprintf("%s %s", key, payload)
}

// EncryptMessage encrypts a message with the provided daily key and message key.
func (e *enigma) EncryptMessage(msg []Char, dk DailyKey, mk MessageKey) EncryptedMessage {
	e.SetDailyKey(dk)
	emk := mk.Encrypt(e)

	e.SetMessageKey(mk)
	payload := e.EncryptDecrypt(msg)

	return EncryptedMessage{
		EncryptedMessageKey: emk,
		Payload:             payload,
	}
}

// DecryptMessage decrypts a message with the provided daily key.
func (e *enigma) DecryptMessage(in EncryptedMessage, dk DailyKey) []Char {
	e.SetDailyKey(dk)
	mk := in.EncryptedMessageKey.Decrypt(e)

	e.SetMessageKey(mk)
	payload := e.EncryptDecrypt(in.Payload)

	return payload
}

// EncryptDecrypt encrypts or decrypts a slice of chars.
func (e *enigma) EncryptDecrypt(chars []Char) []Char {
	res := make([]Char, 0, len(chars))
	for _, c := range chars {
		cc := e.encryptDecryptChar(c)
		res = append(res, cc)
	}
	return res
}

// encryptDecryptChar encrypts or decrypts a single char.
func (e *enigma) encryptDecryptChar(c Char) Char {
	c = e.plugboard.Handle(c)
	c = e.spindle.Handle(c)
	c = e.plugboard.Handle(c)
	return c
}
