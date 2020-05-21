package crypto

import (
	"crypto/rsa"
	"crypto/x509"
)

func EncodePrivate(private *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(private)
}
