package dao

import "fmt"

func UsernameExist(username string) error {
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

func UsernameAndPasswordExist(username string, password string) error {
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

func AddUser(username string, password string) { //
	db, _ := InitDB()
	sqlStr := "insert into userdata values(null, ?, ?)"
	_, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdatePassword(username string, newPassword string) {
	db, _ := InitDB()
	sqlStr := "update userdata set password=? where username=?"
	_, err := db.Exec(sqlStr, newPassword, username)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
	fmt.Println("密码修改成功")
}
