package api

import (
	"fmt"
	"message-board/dao"
)

func SendComment(username string) {
	r := ""
	m := ""
	for {
		fmt.Print("给哪个用户留言:")
		fmt.Scan(&r)
		if !UsernameIsValid(r) {
			fmt.Println("名字过长")
			continue
		}
		break
	}

	for {
		fmt.Print("请输入留言内容：")
		fmt.Scan(&m)
		if !TextIsValid(m) {
			fmt.Println("信息量过大")
			continue
		}
		break
	}
	dao.AddMessage(username, r, m)
}
