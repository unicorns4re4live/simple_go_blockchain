package blockchain

import (
	"bytes"
	"../utils"
	"../crypto"
	"../models"
	"../settings"
)

func NewBlock(trans *models.Transaction, prev_hash []byte) *models.Block {
	var curr_hash = crypto.HashSum(bytes.Join(
		[][]byte{trans.Hash, []byte(settings.User.Hash), utils.ToBytes(settings.DIFFICULTY), prev_hash},
		[]byte{},
	))
	return &models.Block{
		Transaction: *trans,
		Miner: settings.User.Hash,
		Difficulty: settings.DIFFICULTY,
		PrevHash: prev_hash,
		CurrHash: curr_hash,
		Nonce: ProofOfWork(curr_hash),
	}
}
