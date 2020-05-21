package blockchain

import (
	"encoding/hex"
	"../utils"
	"../settings"
)

func CopyBranch(root_branch, new_branch string, last_hash []byte) {
	rows, err := settings.BlockChain.Query(
		"SELECT Hash, Block FROM BlockChain WHERE Branch=$1 ORDER BY Id",
		root_branch,
	)
	utils.CheckError(err)

	var (
		hash = hex.EncodeToString(last_hash)
		hash_str string
		block_str string

		hashes []string
		blocks []string
	)

	for rows.Next() {
		rows.Scan(&hash_str, &block_str)

		hashes = append(hashes, hash_str)
		blocks = append(blocks, block_str)

		if hash == hash_str {
			break
		}
	}
	rows.Close()

	for index := range blocks {
		_, err := settings.BlockChain.Exec(
			"INSERT INTO BlockChain (Branch, Hash, Block) VALUES ($1, $2, $3)",
			new_branch,
			hashes[index],
			blocks[index],
		)
		utils.CheckError(err)
	}
}
