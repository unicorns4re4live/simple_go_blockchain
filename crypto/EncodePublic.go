package crypto

import (
	"crypto/rsa"
	"crypto/x509"
)

func EncodePublic(public *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(public)
}
