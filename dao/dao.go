package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/message_board_database"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
