package settings

import (
	"strings"
	"math/big"
	"encoding/hex"
	"../utils"
	"../crypto"
)

func InitArgs(args []string) {
	var (
		flag_address bool
		flag_branch bool
		has_address bool
		has_branch bool
	)

	for _, arg := range args {
		switch arg {
			case "-a", "--address":
				flag_address = true
				continue
			case "-b", "--branch":
				flag_branch = true
				continue 
		}
		switch {
			case flag_address:
				setAddress(arg)
				has_address = true
				flag_address = false
			case flag_branch:
				setBranch(arg)
				has_branch = true
				flag_branch = false
		}
	}

	if !has_address {
		utils.PrintError("address undefined")
	}

	if !has_branch {
		MainBranch = MASTER_BRANCH
	}

	Target = big.NewInt(1)
	Target.Lsh(Target, 256 - DIFFICULTY)
	LastHash = GetLastHash(MainBranch)

	Authorization()
}

func setAddress(address string) {
	var splited = strings.Split(address, ":")
	if len(splited) != 2 {
		utils.PrintError("len address != 2")
	}
	User.Addr.IPv4 = splited[0]
	User.Addr.Port = ":" + splited[1]
}

func setBranch(branch string) {
	row := BlockChain.QueryRow(
		"SELECT Hash FROM BlockChain WHERE Branch=$1",
		branch,
	)
	var hash_str string
	row.Scan(&hash_str)
	if hash_str == "" {
		utils.PrintError("branch undefined")
	}
	MainBranch = branch
}

func Authorization() {
	if !privateIsExist() {
		createPrivate()
	}
	initAccount()
}

func privateIsExist() bool {
	row := UserData.QueryRow("SELECT Private FROM UserData")
	var str_private string
	row.Scan(&str_private)
	if str_private == "" {
		return false
	}
	return true
}

func createPrivate() {
	private := crypto.GeneratePrivate(2048)
	public := &private.PublicKey
	hashpublic := crypto.PublicToHash(public)

	_, err := UserData.Exec(
		"INSERT INTO UserData (Hash, Private) VALUES ($1, $2)",
		hex.EncodeToString(hashpublic),
		hex.EncodeToString(crypto.EncodePrivate(private)),
	)
	utils.CheckError(err)
}

func initAccount() {
	row := UserData.QueryRow("SELECT Hash, Private FROM UserData")
	var hash_str, private_str string
	row.Scan(&hash_str, &private_str)
	private_bytes, err := hex.DecodeString(private_str)
	utils.CheckError(err)
	User.Hash = hash_str
	User.Keys.Private = crypto.DecodePrivate(private_bytes)
	User.Keys.Public = &(User.Keys.Private).PublicKey
}
