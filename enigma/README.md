# enigma

[![Go Reference](https://pkg.go.dev/badge/github.com/vmorsell/enigma/enigma.svg)](https://pkg.go.dev/github.com/vmorsell/enigma/enigma)

Enigma implementation in Go.

## Example usage

```go
// The daily key needs to be known both when
// encrypting and decrypting.
dk := enigma.NewRandomDailyKey()

// Encrypt.
e1 := enigma.NewEnigma(dk)

payload, _ := enigma.StringToChars("HELLOWORLD")
mk := enigma.NewRandomMessageKey()
msg := e1.EncryptMessage(payload, dk, mk)
fmt.Println(msg)

// Decrypt.
e2 := enigma.NewEnigma(dk)
e2.SetMessageKey(mk)

chars := e2.DecryptMessage(msg, dk)
str := enigma.CharsToString(chars)
fmt.Println(str)
```
