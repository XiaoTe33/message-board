package main

import (
	"fmt"
	"message-board/service"
)

func main() {
	s := "啊啊啊啊啊啊啊啊啊啊哦哦哦哦哦哦哦哦哦啊啊"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	fmt.Println(service.PasswordIsValid(s))
	fmt.Println(service.TextIsValid(s))
}
