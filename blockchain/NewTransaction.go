package blockchain

import (
	"time"
	"bytes"
	"../utils"
	"../crypto"
	"../models"
	"../settings"
)

func NewTransaction(branch, from, to string, value uint64) *models.Transaction {
	var curr_time = time.Now().Format(time.RFC1123)
	var hash = crypto.HashSum(bytes.Join(
		[][]byte{[]byte(branch), []byte(from), []byte(to), utils.ToBytes(value), []byte(curr_time)},
		[]byte{},
	))
	if from != settings.GENESIS_STRING {
		var sum = SumTransaction(branch, from)
		if sum < value {
			return nil
		}
	}
	return &models.Transaction{
		Hash: hash,
		From: from,
		To: to,
		Value: value,
		Timestamp: curr_time,
	}
}
