package main

import (
	"fmt"
	"message-board/dao"
)

func main() {
	_, err := dao.InitDB()
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}
