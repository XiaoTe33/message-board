package dao

import (
	"database/sql"
	"fmt"
	"message-board/model"
	"strconv"
	"time"
)

func FindCommentByMID(MID string) error {
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

func FindMessageByReceiver(receiver string) interface{} {
	db, _ := InitDB()
	sqlStr := "select mid, sender, receiver, time, text from messages where receiver = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare err")
		return nil
	}
	rows, err := stmt.Query(receiver)
	if err != nil {
		fmt.Println("Query err")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Close err")
		}
	}(rows)
	fmt.Printf("%-8v|  %-20v|  %-20v|  %-20v|  %v\n", "MID", "Sender", "Receiver", "Time", "Text")
	var messages = map[string]interface{}{}
	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.MID, &m.Sender, &m.Receiver, &m.Time, &m.Text)
		if err != nil {
			fmt.Println("Scan err")
			return nil
		}

		fmt.Printf("%-8v|  %-20v|  %-20v|  %-20v|  %v\n", m.MID, m.Sender, m.Receiver, m.Time, m.Text)
		messages[m.MID] = m
	}
	return messages
}

func AddMessage(username string, receiver string, text string) {
	db, _ := InitDB()
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into messages(mid, sender, receiver, time, text) values(null, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, receiver, t, text)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func UpdateMessageByMID(mid string, text string) {
	db, _ := InitDB()
	sqlStr := "update messages set text = ? where mid =?"
	_, err := db.Exec(sqlStr, text, mid)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func MIDExist(mid string) bool {
	db, _ := InitDB()
	sqlStr := "select mid from messages where mid = ?"
	stmt, _ := db.Prepare(sqlStr)
	row := stmt.QueryRow(mid)
	id := ""
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func DeleteMessageByMID(mid string) {
	db, _ := InitDB()
	sqlStr := "update messages set deleted = '留言已删除' where mid =?"
	_, err := db.Exec(sqlStr, mid)
	if err != nil {
		fmt.Println(err)
		return
	}
}
