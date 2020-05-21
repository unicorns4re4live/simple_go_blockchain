package blockchain

import (
	"../models"
	"../settings"
)

func HardFork(root_branch, new_branch string, last_hash []byte) *models.Block {
	CopyBranch(root_branch, new_branch, last_hash)
	var trans = NewTransaction(new_branch, settings.User.Hash, settings.User.Hash, 0)
	var block = NewBlock(trans, last_hash)
	return block
}
