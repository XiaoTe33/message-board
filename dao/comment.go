package dao

import (
	"database/sql"
	"fmt"
	"message-board/model"
	"strconv"
	"time"
)

func FindCommentByMID(mid string) map[string]interface{} {
	db, _ := InitDB()
	sqlStr := "select mid, cid, sender, time, text from comments where mid = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare err")
		return nil
	}
	rows, err := stmt.Query(mid)
	if err != nil {
		fmt.Println("query err")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	var comments = map[string]interface{}{}
	for rows.Next() {
		var c model.Comment
		err := rows.Scan(&c.MID, &c.CID, &c.Sender, &c.Time, &c.Text)
		if err != nil {
			fmt.Println("scan err")
			return nil
		}
		comments[c.CID] = c
		return comments
	}
	return nil
}

func AddComment(sender string, mid string, text string) {
	db, _ := InitDB()
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into comments(mid, sender, time, text) values(?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, mid, sender, t, text)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func CIDExist(cid string) bool {
	db, _ := InitDB()
	sqlStr := "select cid from comments where cid = ?"
	stmt, _ := db.Prepare(sqlStr)
	row := stmt.QueryRow(cid)
	id := ""
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func RIDExist(rid string) bool {
	return CIDExist(rid)
}

func UpdateCommentByCID(cid string, text string) {
	db, _ := InitDB()
	sqlStr := "update comments set text = ? where cid =?"
	_, err := db.Exec(sqlStr, text, cid)
	if err != nil {
		fmt.Println("exec err")
		return
	}
}

func DeleteCommentByCID(cid string) {
	db, _ := InitDB()
	sqlStr := "update comments set deleted = '留言已删除' where cid =?"
	_, err := db.Exec(sqlStr, cid)
	if err != nil {
		fmt.Println("exec err")
		return
	}
}

func AddResponseComment(sender string, mid string, rid string, text string) {
	db, _ := InitDB()
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into comments(mid, sender, time, text, rid) values(?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, mid, sender, t, text, rid)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}
