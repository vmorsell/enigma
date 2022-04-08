package enigma

import "fmt"

const (
	payloadGroupSize = 5
)

// Enigma represents an Enigma instance.
type Enigma interface {
	EncryptMessage(payload []Char, dk DailyKey, mk MessageKey) EncryptedMessage
	DecryptMessage(msg EncryptedMessage, dk DailyKey) []Char
	EncryptDecrypt(chars []Char) []Char
	SetDailyKey(key DailyKey)
	SetMessageKey(key MessageKey)
}

// enigma implements the Enigma logic.
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

type EncryptedMessage struct {
	EncryptedMessageKey MessageKey
	Payload             []Char
}

func (msg EncryptedMessage) String() string {
	key := fmt.Sprintf("%s %s", charsToString(msg.EncryptedMessageKey.Positions[:3]), charsToString(msg.EncryptedMessageKey.Positions[3:]))

	payload := charsToString(msg.Payload)
	for i := payloadGroupSize; i < len(payload); i += payloadGroupSize + 1 {
		payload = fmt.Sprintf("%s %s", payload[:i], payload[i:])
	}

	return fmt.Sprintf("%s %s", key, payload)
}

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
