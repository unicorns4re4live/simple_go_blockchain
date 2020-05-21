package network

import (
	"os"
	"fmt"
	"time"
	"strings"
	"math/rand"
	"encoding/json"
	"encoding/base64"
	"../utils"
	"../models"
	"../settings"
	"../blockchain"
)

func ClientTCP() {
	var message string
	var splited []string
	for {
		message = utils.InputString()
		splited = strings.Split(message, " ")
		switch splited[0] {
			case settings.CMD_EXIT: os.Exit(0)
			case settings.CMD_WHOAMI: whoami()
			case settings.CMD_GENESIS: genesis()
			case settings.CMD_NETWORK: network()
			case settings.CMD_BRANCH: branch()
			case settings.CMD_BALANCE: balance(splited[1:])
			case settings.CMD_CONNECT: connectTo(splited[1:])
			case settings.CMD_HARDFORK: hardfork(splited[1:])
			case settings.CMD_SELECT_BRANCH: selectBranch(splited[1:])
			case settings.CMD_REQ_PUSH: reqPushBlock(splited[1:])
			case settings.CMD_READ_BLOCKS: readBlocks()
			case settings.CMD_DOWNLOAD_BRANCH: downloadBranch(splited[1:])
		}
	}
}

func downloadBranch(splited []string) {
	if len(splited) != 1 {
		fmt.Println("len splited != 1")
		return
	}
	var new_pack = models.PackageTCP {
		To: models.To {
			Addr: randMiner(),
		},
		Head: models.Head {
			Title: settings.TITLE_CONNECT,
			Mode: settings.MODE_SAVE,
		},
		Body: models.Body {
			Branch: splited[0],
		},
	}
	SendPackage(&new_pack)
}

func branch() {
	fmt.Println("Branch:", settings.MainBranch)
}

func selectBranch(splited []string) {
	if len(splited) != 1 {
		fmt.Println("len splited != 1")
		return
	}
	var selected_branch = splited[0]
	var last = settings.GetLastHash(selected_branch)
	if last == nil {
		fmt.Println("branch undefined")
		return
	}
	settings.LastHash = last
	settings.MainBranch = selected_branch
}

func hardfork(splited []string) {
	if len(splited) != 2 {
		fmt.Println("len splited != 2")
		return
	}

	var (
		new_branch = splited[0]
		blockhash = splited[1]
	)

	decoded, err := base64.StdEncoding.DecodeString(blockhash)
	if err != nil {
		fmt.Println("incorrected base64 hash")
		return
	}
	var block = blockchain.HardFork(settings.MainBranch, new_branch, decoded)
	blockchain.PushBlock(new_branch, block)

	block_json, err := json.MarshalIndent(*block, "", "\t")
	utils.CheckError(err)

	var new_pack = models.PackageTCP {
		Head: models.Head {
			Title: settings.TITLE_BLOCK,
			Mode: settings.MODE_HARD,
		},
		Body: models.Body {
			Data: string(block_json),
			Branch: settings.MainBranch + settings.SEPARATOR + new_branch,
		},
	}

	SendGlobalPackage(&new_pack)
}

func readBlocks() {
	var (
		id uint64
		block_str string
	)
	rows, err := settings.BlockChain.Query(
		"SELECT Block FROM BlockChain WHERE Branch=$1",
		settings.MainBranch,
	)
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&block_str)
		fmt.Printf("Block: %d\n%s\n", id, block_str)
		id++
	}
}

func network() {
	for addr := range settings.Connections {
		fmt.Println("|", addr)
	}
}

func randMiner() string {
	var length = len(settings.Connections)
	if length == 0 { return "" }
	rand.Seed(time.Now().UnixNano())
	num := rand.Int() % length
	for addr := range settings.Connections {
		if num == 0 {
			return addr
		}
		num--
	}
	return ""
}

func reqPushBlock(splited []string) {
	if len(splited) != 2 {
		fmt.Println("len splited != 2")
		return 
	}
	var new_pack = models.PackageTCP {
		To: models.To {
			Addr: randMiner(),
		},
		Head: models.Head {
			Title: settings.TITLE_BLOCK,
			Mode: settings.MODE_PUSH,
		},
		Body: models.Body {
			Data: splited[0] + settings.SEPARATOR + splited[1],
			Branch: settings.MainBranch,
		},
	}
	SendPackage(&new_pack)
}

func connectTo(splited []string) {
	if len(splited) != 1 {
		fmt.Println("len splited != 1")
		return
	}
	var address = splited[0]
	settings.Connections[address] = true
	var new_pack = models.PackageTCP {
		To: models.To {
			Addr: address,
		},
		Head: models.Head {
			Title: settings.TITLE_CONNECT,
			Mode: settings.MODE_SAVE,
		},
		Body: models.Body {
			Branch: settings.MainBranch,
		},
	}
	SendPackage(&new_pack)
}

func whoami() {
	fmt.Println("Whoami:", settings.User.Hash)
}

func genesis() {
	var block = blockchain.CreateGenesis(settings.User.Hash)
	if block == nil {
		fmt.Println("genesis is exist")
		return
	}
	blockchain.PushBlock(settings.MASTER_BRANCH, block)
	var new_pack = convertBlockToPackage(settings.MASTER_BRANCH, block)
	SendGlobalPackage(new_pack)
}

func convertBlockToPackage(branch string, block *models.Block) *models.PackageTCP {
	block_json, err := json.MarshalIndent(*block, "", "\t")
	utils.CheckError(err)
	return &models.PackageTCP{
		Head: models.Head {
			Title: settings.TITLE_BLOCK,
			Mode: settings.MODE_SAVE,
		},
		Body: models.Body {
			Data: string(block_json),
			Branch: branch,
		},
	}
}

func balance(splited []string) {
	var account string
	if len(splited) == 0 {
		account = settings.User.Hash
	} else {
		account = splited[0]
	}
	fmt.Println("Balance:", blockchain.SumTransaction(settings.MainBranch, account), "coins.")
}
