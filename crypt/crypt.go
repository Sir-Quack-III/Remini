package crypt

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSHA256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))

	bs := h.Sum(nil)

	return string(bs)
}

func EncryptAES(plaintext string, k string) string {
	// create cipher
	key := []byte(k)
	c, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
