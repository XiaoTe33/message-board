package dao

import (
	"database/sql"
	"fmt"
	"message-board/model"
	"strconv"
	"time"
)

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
func AddUser(username string, password string) { //
	fmt.Println(username, password)
}

func FindComment(MID int) error {
	db, _ := InitDB()
	sqlStr := "select mid, cid, sender, receiver, time, text from comments where mid = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	rows, err := stmt.Query(MID)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	fmt.Printf("%-8v|  %-8v|  %-20v|  %-20v|  %-20v|  %v\n", "MID", "CID", "Sender", "Receiver", "Time", "Text")
	for rows.Next() {
		var c model.Comment
		err := rows.Scan(&c.MID, &c.CID, &c.Sender, &c.Receiver, &c.Time, &c.Text)
		if err != nil {
			return err
		}

		fmt.Printf("%-8v|  %-8v|  %-20v|  %-20v|  %-20v|  %v\n", c.MID, c.CID, c.Sender, c.Receiver, c.Time, c.Text)
	}
	return nil
}

func FindMessage(receiver string) error {
	db, _ := InitDB()
	sqlStr := "select id, sender, receiver, time, text from messages where receiver = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	rows, err := stmt.Query(receiver)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	fmt.Printf("%-8v|  %-20v|  %-20v|  %-20v|  %v\n", "ID", "Sender", "Receiver", "Time", "Text")
	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.ID, &m.Sender, &m.Receiver, &m.Time, &m.Text)
		if err != nil {
			return err
		}

		fmt.Printf("%-8v|  %-20v|  %-20v|  %-20v|  %v\n", m.ID, m.Sender, m.Receiver, m.Time, m.Text)
	}
	return nil
}

func UpdateP(username string, password string) {
	db, _ := InitDB()
	sqlStr := "update userdata set password=? where username=?"
	_, err := db.Exec(sqlStr, password, username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("密码修改成功")
}

func AddMessage(username string, receiver string, text string) {
	db, _ := InitDB()
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into messages values(null, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, receiver, t, text)
	if err != nil {
		fmt.Println(err)
		return
	}
}
