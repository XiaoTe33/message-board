package main

import (
	"fmt"
	"message-board/dao"
)

func main() {
	a := dao.FindDialogByCID("9")
	fmt.Println(a)
}
