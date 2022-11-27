package main

import (
	"fmt"
	"message-board/dao"
	"message-board/test/test05"
)

func main() {
	//api.ChangeMessage()
	main04()
}
func main06() {
	//api.SendMessage("xiaote")
	main03()
}
func main05() {
	test05.ChangeP()
}
func main04() {
	//u := "xiaote33"
	//err := dao.FindMessageByReceiver(u)
	ID := "1"
	err := dao.FindCommentByMID(ID)
	if err != nil {
		fmt.Println(err)
	}
}

func main03() {
	test05.Register()
	//api.Login()
}
func main02() {
	u := "xiaote"
	p := "1"
	err := dao.UsernameAndPasswordExist(u, p)
	if err != nil {
		fmt.Println(err)
	}
}
func main01() {
	u := "xiaote"
	err := dao.UsernameExist(u)
	if err != nil {
		fmt.Println(err)
	}
}
