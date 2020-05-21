package blockchain

import (
	"fmt"
	"math"
	"bytes"
	"math/big"
	"../utils"
	"../crypto"
	"../settings"
)

func ProofOfWork(hash []byte) uint64 {
	var (
		new_hash []byte
		nonce uint64
		number big.Int
	)
	for nonce < math.MaxUint64 {
		new_hash = crypto.HashSum(bytes.Join(
			[][]byte{utils.ToBytes(nonce), hash},
			[]byte{},
		))
		if settings.VIEW_MINING {
			fmt.Printf("\r%x", new_hash)
		}
		number.SetBytes(new_hash)
		if number.Cmp(settings.Target) == -1 {
			fmt.Println()
			break
		}
		nonce++
	}
	return nonce
}
