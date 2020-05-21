package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"../utils"
)

func GeneratePrivate(bits int) *rsa.PrivateKey {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	utils.CheckError(err)
	return private
}
