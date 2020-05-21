package crypto

import (
	"crypto/rsa"
)

func PublicToHash(public *rsa.PublicKey) []byte {
	return HashSum(EncodePublic(public))
}
