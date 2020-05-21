package utils

import (
	"os"
	"fmt"
)

func PrintError(err string) {
	fmt.Println("[ERROR]:", err)
	os.Exit(1)
}
