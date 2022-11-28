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

func FindDialogByCID(cid string) map[string]model.Comment {
	var dialog = map[string]model.Comment{}
	db, _ := InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	sqlStr := "select cid from comments "
	rows, err := db.Query(sqlStr)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("close err")
		}
	}(rows)
	if err != nil {
		fmt.Println("query err")
		return nil
	}
	for rows.Next() {
		cid2 := ""
		err := rows.Scan(&cid2)
		if err != nil {
			fmt.Println("Scan err")
			return nil
		}
		if FindFatherCommentByCID(cid) == FindFatherCommentByCID(cid2) {
			db2, _ := InitDB()
			sqlStr2 := "select mid, cid, sender, time, text, deleted, rid from comments where cid = ?"
			row := db2.QueryRow(sqlStr2, cid2)
			var comment model.Comment
			err := row.Scan(&comment.MID, &comment.CID, &comment.Sender,
				&comment.Time, &comment.Text, &comment.Deleted, &comment.RID)
			if err != nil {
				fmt.Println("Scan err2")
				return nil
			}
			dialog[comment.CID] = comment
			err2 := db2.Close()
			if err2 != nil {
				return nil
			}
		}

	}
	return dialog
}
func FindFatherCommentByCID(cid string) string {
	db, _ := InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	//1.找到根节点cid
	for {
		sqlStr := "select rid from comments where cid =?"
		row := db.QueryRow(sqlStr, cid)
		rid := ""
		err := row.Scan(&rid)
		if err != nil {
			fmt.Println("scan err")
			return ""
		}
		if rid == "0" {
			break
		}
		cid = rid
	}
	return cid
}
