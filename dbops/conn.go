package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbConn, err = sql.Open("mysql", "root:@tcp(localhost:3306)/mysql?charset=utf8") //我也不知道密码是多少
	if err != nil {
		panic(err.Error())
	}
}

var (
	dbConn *sql.DB
	err    error
)
