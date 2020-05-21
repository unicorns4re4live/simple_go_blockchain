package blockchain

import (
	"bytes"
	"../models"
)

func ValidateChain(blocks []models.Block) bool {
	var length = len(blocks)
	if length <= 1 {
		return true 
	}

	for i := 1; i < length; i++ {
		if !bytes.Equal(blocks[i-1].CurrHash, blocks[i].PrevHash) {
			return false
		}
	}
	return true
}
