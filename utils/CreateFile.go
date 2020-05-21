package utils

import (
	"os"
)

func CreateFile(name string) {
	file, err := os.Create(name)
	CheckError(err)
	file.Close()
}
