package utils

import (
	"bytes"
	"encoding/binary"
)

func ToBytes(num uint64) []byte {
	var data = new(bytes.Buffer)
	err := binary.Write(data, binary.BigEndian, num)
	CheckError(err)
	return data.Bytes()
}
