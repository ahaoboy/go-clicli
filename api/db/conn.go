package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:changhao2333@tcp(localhost:3306)/clicli?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

//root:changhao2333@tcp(localhost:3306)/clicli?charset=utf8