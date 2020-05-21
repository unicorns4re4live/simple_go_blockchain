package utils

import (
	"os"
	"bufio"
	"strings"
)

func InputString() string {
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(data, "\n", "", -1)
}
