package models

import (
	"crypto/rsa"
)

type Keys struct {
	Private *rsa.PrivateKey
	Public *rsa.PublicKey
}

type Addr struct {
	IPv4 string
	Port string
}

type User struct {
	Hash string
	Keys Keys
	Addr Addr
}
