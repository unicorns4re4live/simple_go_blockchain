package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"../utils"
)

func DecodePrivate(private []byte) *rsa.PrivateKey {
	priv, err := x509.ParsePKCS1PrivateKey(private)
	utils.CheckError(err)
	return priv
}
