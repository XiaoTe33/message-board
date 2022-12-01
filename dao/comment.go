package dao

import (
	"database/sql"
	"fmt"
	"message-board/model"
	"strconv"
	"time"
)

func FindCommentByMID(mid string) map[string]interface{} {
	sqlStr := "select mid, cid, sender, time, text from comments where mid = ?"
	stmt, err := Db.Prepare(sqlStr)
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
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into comments(mid, sender, time, text) values(?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, mid, sender, t, text)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func FindCID(cid string) bool {
	sqlStr := "select cid from comments where cid = ?"
	stmt, _ := Db.Prepare(sqlStr)
	row := stmt.QueryRow(cid)
	id := ""
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func FindRID(rid string) bool {
	return FindCID(rid)
}

func UpdateCommentByCID(cid string, text string) {
	sqlStr := "update comments set text = ? where cid =?"
	_, err := Db.Exec(sqlStr, text, cid)
	if err != nil {
		fmt.Println("exec err")
		return
	}
}

func DeleteCommentByCID(cid string) {

	sqlStr := "update comments set deleted = '留言已删除' where cid =?"
	_, err := Db.Exec(sqlStr, cid)
	if err != nil {
		fmt.Println("exec err")
		return
	}
}

func AddResponseComment(sender string, mid string, rid string, text string) {
	now := time.Now()
	t := strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sqlStr := "insert into comments(mid, sender, time, text, rid) values(?, ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, mid, sender, t, text, rid)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func FindDialogByCID(cid string) map[string]model.Comment {
	//找到传入评论的根节点
	for {
		sqlStr := "select rid from comments where cid =?"
		row := Db.QueryRow(sqlStr, cid)
		rid := ""
		err := row.Scan(&rid)
		if err != nil {
			fmt.Println("scan err")
			return nil
		}
		if rid == "0" {
			break
		}
		cid = rid
	}
	sqlStr := "select mid, cid, sender, time, text, deleted, rid from comments where cid =?"
	var rt model.Comment
	row := Db.QueryRow(sqlStr, cid)
	err := row.Scan(&rt.MID, &rt.CID, &rt.Sender, &rt.Time, &rt.Text, &rt.Deleted, &rt.RID)
	if err != nil {
		fmt.Println("scan err")
		return nil
	}
	var dialogs = map[string]model.Comment{cid: rt}

	FindChildren(cid, dialogs)
	return dialogs
}

// FindChildren 深搜遍历
func FindChildren(cid string, temp map[string]model.Comment) {
	sqlStr := "select mid, cid, sender, time, text, deleted, rid from comments where rid =?"
	rows, err2 := Db.Query(sqlStr, cid)
	if err2 != nil {
		fmt.Println("query err")
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("close err")
		}
	}(rows)
	for rows.Next() {
		var c model.Comment
		err := rows.Scan(&c.MID, &c.CID, &c.Sender, &c.Time, &c.Text, &c.Deleted, &c.RID)
		if err != nil {
			fmt.Println("scan err")
		}
		temp[c.CID] = c
		FindChildren(c.CID, temp)
	}
}
