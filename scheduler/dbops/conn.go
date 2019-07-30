// Here we directly copy the code from api_conn, since we use the same database
// However, in the real case, we want resources in the distributed form in case of influencing each other
// therefore, we choose to use different datacases in different service.
package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:@tcp(localhost:3306)/mysql?charset=utf8") //我也不知道密码是多少
	if err != nil {
		panic(err.Error())
	}
}
