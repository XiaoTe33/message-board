package dao

import "fmt"

func UExist(username string) error {
	db, _ := InitDB()
	sqlStr := "select username from userdata where username = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(username)
	u := ""
	err = row.Scan(&u)
	if err != nil {
		return err
	}
	return nil
}

func UPExist(username string, password string) error {
	db, _ := InitDB()
	sqlStr := "select username from userdata where username = ? and password = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(username, password)
	u := ""
	err = row.Scan(&u)
	if err != nil {
		return err
	}
	return nil
}
func AddUser(username string, password string) {
	fmt.Println(username, password)
}
