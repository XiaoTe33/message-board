package test05

import (
	"fmt"
	"message-board/dao"
	"message-board/service"
)

func SendMessage(username string) {
	r := ""
	m := ""
	for {
		fmt.Print("给哪个用户留言:")
		fmt.Scan(&r)
		if !service.UsernameIsValid(r) {
			fmt.Println("名字过长")
			continue
		}
		break
	}

	for {
		fmt.Print("请输入留言内容：")
		fmt.Scan(&m)
		if !service.TextIsValid(m) {
			fmt.Println("信息量过大")
			continue
		}
		break
	}

}

func ChangeMessageByMID() {
	for {
		mid := ""
		fmt.Print("请输入留言mid:")
		fmt.Scan(&mid)
		text := ""
		fmt.Print("请输入修改后的文本:")
		fmt.Scan(&text)
		if !dao.MIDExist(mid) {
			fmt.Println("mid不存在")
			continue
		}
		if !service.TextIsValid(text) {
			fmt.Println("文本太长")
			continue
		}
		dao.UpdateMessageByMID(mid, text)
		fmt.Println("修改成功")
		break
	}

}
