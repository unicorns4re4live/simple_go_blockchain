package blockchain

import (
	"../models"
	"../settings"
)

func CreateGenesis(account string) *models.Block {
	if GenesisIsExist() {
		return nil
	}
	var trans = NewTransaction(settings.MASTER_BRANCH, settings.GENESIS_STRING, account, 100)
	return NewBlock(trans, []byte{})
}
