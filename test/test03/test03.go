package main

import (
	"fmt"
	"message-board/api"
	"message-board/dao"
)

func main() {
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
