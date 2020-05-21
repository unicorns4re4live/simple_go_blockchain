package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"../utils"
	"../models"
	"../settings"
)

func PushBlock(branch string, block *models.Block) {
	block_json, err := json.MarshalIndent(*block, "", "\t")
	utils.CheckError(err)
	_, err = settings.BlockChain.Exec(
		"INSERT INTO BlockChain (Branch, Hash, Block) VALUES ($1, $2, $3)",
		branch,
		hex.EncodeToString(block.CurrHash),
		string(block_json),
	)
	utils.CheckError(err)
}
