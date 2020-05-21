package main

import (
	"os"
	"./network"
	"./settings"
)

func main() {
	settings.OpenUserData("userdata.db")
	settings.OpenBlockChain("blockchain.db")
	settings.InitArgs(os.Args)
	go network.ServerTCP()
	network.ClientTCP()
}
