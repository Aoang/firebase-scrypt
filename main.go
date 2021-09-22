// Package scrypt implements the modified Scrypt algorithm used by Firebase Auth.
package scrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

const (
	P      = 1
	KeyLen = 32
)

type App struct {
	SignerKey     []byte
	SaltSeparator []byte
	Rounds        int
	MemCost       int
	P             int
	KeyLen        int
}

func New(signerKey, saltSeparator string, rounds, memCost int) (*App, error) {
	var (
		app = &App{
			Rounds:  rounds,
			MemCost: memCost,
			P:       P,
			KeyLen:  KeyLen,
		}
		err error
	)

	if app.SignerKey, err = base64.StdEncoding.DecodeString(signerKey); err != nil {
		return nil, err
	}
	if app.SaltSeparator, err = base64.StdEncoding.DecodeString(saltSeparator); err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) EncodeToString(password, salt []byte) (string, error) {
	res, err := a.Encode(password, salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

func (a *App) Encode(password, salt []byte) ([]byte, error) {
	return key(password, salt, a.SignerKey, a.SaltSeparator, a.Rounds, a.MemCost, a.P, a.KeyLen)
}

func (a *App) Verify(password, salt []byte, passwordHash string) bool {
	h, err := a.EncodeToString(password, salt)
	if err != nil {
		return false
	}

	return h == passwordHash
}

func (a *App) FirebaseVerify(password, salt, passwordHash string) bool {
	_salt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}

	var hs string
	if hs, err = a.EncodeToString([]byte(password), _salt); err != nil {
		return false
	}
	return hs == passwordHash
}

func Key(password, salt []byte, signerKey, saltSeparator string, rounds, memCost, p, keyLen int) ([]byte, error) {
	var (
		sk, ss []byte
		err    error
	)

	if sk, err = base64.StdEncoding.DecodeString(signerKey); err != nil {
		return nil, err
	}
	if ss, err = base64.StdEncoding.DecodeString(saltSeparator); err != nil {
		return nil, err
	}

	return key(password, salt, sk, ss, rounds, memCost, p, keyLen)
}

func key(password, salt, signerKey, saltSeparator []byte, rounds, memCost, p, keyLen int) ([]byte, error) {
	ck, err := scrypt.Key(password, append(salt, saltSeparator...), 1<<memCost, rounds, p, keyLen)
	if err != nil {
		return nil, err
	}

	var block cipher.Block
	if block, err = aes.NewCipher(ck); err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(signerKey))
	stream := cipher.NewCTR(block, cipherText[:aes.BlockSize])
	stream.XORKeyStream(cipherText[aes.BlockSize:], signerKey)
	return cipherText[aes.BlockSize:], nil
}
