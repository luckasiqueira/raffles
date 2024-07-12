package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"raffles/utils/info"
	"raffles/utils/logger"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", info.Env["DB_USER"], info.Env["DB_PASS"], info.Env["DB_HOST"], info.Env["DB_PORT"], info.Env["DB_NAME"])

// Connect starts a database connection using .env data
func Connect() *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.NewLog(err.Error())
	}
	return db
}
