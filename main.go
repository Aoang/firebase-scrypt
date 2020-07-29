// Package scrypt implements the modified Scrypt algorithm used by Firebase Auth.
package scrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
)

var Default *crypt

// avoid panic caused by null pointers.
func init() {
	Default = &crypt{}
}

func Encode(password, salt string) (string, error) {
	return Default.Encode(password, salt)
}

func Verify(password, passwordHash, salt string) bool {
	return Default.Verify(password, passwordHash, salt)
}

const (
	P      = 1
	KeyLen = 32
)

type crypt struct {
	SignerKey     []byte
	SaltSeparator []byte
	Rounds        int
	MemCost       int
	P             int
	KeyLen        int
}

func New(signerKey, saltSeparator string, rounds, memCost int) *crypt {
	sk, _ := base64.StdEncoding.DecodeString(signerKey)
	ss, _ := base64.StdEncoding.DecodeString(saltSeparator)

	return &crypt{
		SignerKey:     sk,
		SaltSeparator: ss,
		Rounds:        rounds,
		MemCost:       memCost,
		P:             P,
		KeyLen:        KeyLen,
	}
}

func (c *crypt) Encode(password, salt string) (string, error) {
	if c.SaltSeparator == nil || c.SignerKey == nil {
		return "", errors.New("config error")
	}

	s, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	ck, err := scrypt.Key([]byte(password),
		append(s, c.SaltSeparator...),
		1<<c.MemCost, c.Rounds, P, KeyLen)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(ck)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(c.SignerKey))

	stream := cipher.NewCTR(block, cipherText[:aes.BlockSize])
	stream.XORKeyStream(cipherText[aes.BlockSize:], c.SignerKey)

	result := base64.StdEncoding.EncodeToString(cipherText[aes.BlockSize:])
	return result, nil
}

func (c *crypt) Verify(password, passwordHash, salt string) bool {
	h, err := c.Encode(password, salt)
	if err != nil {
		return false
	}

	return h == passwordHash
}
