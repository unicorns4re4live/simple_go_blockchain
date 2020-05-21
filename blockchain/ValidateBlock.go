package blockchain

import (
	"bytes"
	"../models"
	"../settings"
)

func ValidateBlock(branch_name string, block models.Block) bool {
	if !bytes.Equal(settings.GetLastHash(branch_name), block.PrevHash) {
		return false
	}
	return true
}
