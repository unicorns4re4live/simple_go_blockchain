package settings

import (
	"math/big"
	"database/sql"
	"../models"
)

var (
	LastHash []byte
	Target *big.Int
	MainBranch string
	BlockChain *sql.DB
	UserData *sql.DB
)

var (
	User models.User
	Connections = make(map[string]bool)
)

const (
	DIFFICULTY = 20
	VIEW_MINING = true
	MASTER_BRANCH = "master"
	GENESIS_STRING = "[GENESIS:STRING]"
	SEPARATOR = "[SEPARATOR]"
)

const (
	CMD_EXIT = ":exit"
	CMD_WHOAMI = ":whoami"
	CMD_NETWORK = ":network"
	CMD_GENESIS = ":genesis"
	CMD_BALANCE = ":balance"
	CMD_CONNECT = ":connect"
	CMD_REQ_PUSH = ":push"
	CMD_BRANCH = ":branch"
	CMD_READ_BLOCKS = ":blocks"
	CMD_HARDFORK = ":hardfork"
	CMD_SELECT_BRANCH = ":select"
	CMD_DOWNLOAD_BRANCH = ":download"
)

const (
	TITLE_CONNECT = "[TITLE:CONNECT]"
	TITLE_BLOCK = "[TITLE:BLOCK]"
	TITLE_ERROR = "[TITLE:ERROR]"
)

const (
	MODE_SAVE = "[MODE:SAVE]"
	MODE_PUSH = "[MODE:PUSH]"
	MODE_HARD = "[MODE:HARD]"
	MODE_READ = "[MODE:READ]"
)
