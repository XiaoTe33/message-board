package main

import (
	"fmt"
	"message-board/api"
)

func main() {
	s := "啊啊啊啊啊啊啊啊啊啊哦哦哦哦哦哦哦哦哦啊啊"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	fmt.Println(api.PasswordIsValid(s))
	fmt.Println(api.TextIsValid(s))
}
