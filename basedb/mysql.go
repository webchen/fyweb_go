package basedb

import (
	// base mysql package
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// wxConnStr the connection string of mysql
var wxConnStr = "root:pass@tcp(ip:3306)/wx?charset=utf8"

// WxDb  connection to the database of "wx"
var WxDb = sqlx.MustConnect("mysql", wxConnStr)

func init() {
	WxDb.SetMaxOpenConns(500)
	WxDb.SetMaxIdleConns(300)
	//WxDb.Ping()
}
