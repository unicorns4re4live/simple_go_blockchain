package models

type Transaction struct {
	Hash []byte
	From string
	To string
	Value uint64
	Timestamp string
}
