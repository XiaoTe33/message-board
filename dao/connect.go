package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (Db *sql.DB, err error) {
	dsn := "root:root@tcp(localhost:3306)/message_board_database"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return Db, err
	}
	err = Db.Ping()
	if err != nil {
		return Db, err
	}
	return Db, nil
}
