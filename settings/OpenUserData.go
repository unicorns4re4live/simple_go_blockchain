package settings

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"../utils"
)

func OpenUserData(name string) {
	var err error

	if !utils.FileIsExist(name) {
		utils.CreateFile(name)
	}

	UserData, err = sql.Open("sqlite3", name)
	utils.CheckError(err)

	_, err = UserData.Exec(`
CREATE TABLE IF NOT EXISTS UserData (
	Hash VARCHAR(32) UNIQUE,
	Private VARCHAR(4096) UNIQUE
);
`)
	utils.CheckError(err)
}
