package dao

import (
	"fmt"
)

func FindUsername(username string) error {
	sqlStr := "select username from userdata where username = ?"
	stmt, err := Db.Prepare(sqlStr)
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

func FindUsernameAndPassword(username string, password string) error {
	sqlStr := "select username from userdata where username = ? and password = ?"
	stmt, err := Db.Prepare(sqlStr)
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

	sqlStr := "insert into userdata(uid, username, password) values(null, ?, ?)"
	_, err := Db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func UpdatePassword(username string, newPassword string) {
	sqlStr := "update userdata set password=? where username=?"
	_, err := Db.Exec(sqlStr, newPassword, username)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
	fmt.Println("密码修改成功")
}
