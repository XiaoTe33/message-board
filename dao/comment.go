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
	sqlStr := "select mid, cid, sender, time, text from comments where mid = ?"
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
	fmt.Printf("%-8v|  %-8v|  %-20v|  %-20v|  %v\n", "MID", "CID", "Sender", "Time", "Text")
	for rows.Next() {
		var c model.Comment
		err := rows.Scan(&c.MID, &c.CID, &c.Sender, &c.Time, &c.Text)
		if err != nil {
			return err
		}

		fmt.Printf("%-8v|  %-8v|  %-20v|  %-20v|  %v\n", c.MID, c.CID, c.Sender, c.Time, c.Text)
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
