package settings

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"../utils"
)

func OpenBlockChain(name string) {
	var err error

	if !utils.FileIsExist(name) {
		utils.CreateFile(name)
	}

	BlockChain, err = sql.Open("sqlite3", name)
	utils.CheckError(err)

	_, err = BlockChain.Exec(`
CREATE TABLE IF NOT EXISTS BlockChain (
	Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	Branch VARCHAR(32),
	Hash VARCHAR(32),
	Block VARCHAR(512)
);
`)
	utils.CheckError(err)
}
