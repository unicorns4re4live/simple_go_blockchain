package blockchain

import (
	"encoding/json"
	"../utils"
	"../models"
	"../settings"
)

func SumTransaction(branch, account string) uint64 {
	var (
		balance uint64
		block models.Block
		block_str string
	)

	rows, err := settings.BlockChain.Query(
		"SELECT Block FROM BlockChain WHERE Branch=$1 ORDER BY Id",
		branch,
	)
	utils.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&block_str)
		json.Unmarshal([]byte(block_str), &block)
		if block.Transaction.From == account {
			balance -= block.Transaction.Value
		}
		if block.Transaction.To == account {
			balance += block.Transaction.Value
		}
	}

	return balance
}
