package models

type Block struct {
	Transaction Transaction
	Miner string
	Difficulty uint8
	PrevHash []byte
	CurrHash []byte
	Nonce uint64
}
