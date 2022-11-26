package main

import (
	"fmt"
	"message-board/api"
	"message-board/dao"
)

func main() {
	api.ChangeP()
}
func main04() {
	//u := "xiaote33"
	//err := dao.FindMessage(u)
	ID := 1
	err := dao.FindComment(ID)
	if err != nil {
		fmt.Println(err)
	}
}

func main03() {
	//api.Register()
	api.Login()
}
func main02() {
	u := "xiaote"
	p := "1"
	err := dao.UPExist(u, p)
	if err != nil {
		fmt.Println(err)
	}
}
func main01() {
	u := "xiaote"
	err := dao.UExist(u)
	if err != nil {
		fmt.Println(err)
	}
}
