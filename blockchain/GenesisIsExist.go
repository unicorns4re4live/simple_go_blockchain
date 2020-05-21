package blockchain

import (
	"../settings"
)

func GenesisIsExist() bool {
	var hash_str string
	row := settings.BlockChain.QueryRow("SELECT Hash FROM BlockChain")
	row.Scan(&hash_str)
	if hash_str == "" {
		return false
	}
	return true
}
