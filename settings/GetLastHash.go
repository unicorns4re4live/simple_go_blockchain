package settings

import (
	"encoding/hex"
)

func GetLastHash(branch string) []byte {
	row := BlockChain.QueryRow(
		"SELECT Hash FROM BlockChain WHERE Branch=$1 ORDER BY Id DESC",
		branch,
	)
	var str_hash string
	row.Scan(&str_hash)
	if str_hash == "" {
		return nil
	}
	hash, err := hex.DecodeString(str_hash)
	if err != nil {
		return nil
	}
	return hash
}
