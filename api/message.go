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

func ChangeMessageByMID() {
	for {
		mid := 0
		fmt.Print("请输入留言mid:")
		fmt.Scan(&mid)
		text := ""
		fmt.Print("请输入修改后的文本:")
		fmt.Scan(&text)
		if !dao.MIDExist(mid) {
			fmt.Println("mid不存在")
			continue
		}
		if !TextIsValid(text) {
			fmt.Println("文本太长")
			continue
		}
		dao.UpdateMessageByMID(mid, text)
		fmt.Println("修改成功")
		break
	}

}
