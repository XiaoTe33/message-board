package service

import (
	"fmt"
	"message-board/util"
	"regexp"
	"strings"
)

func UsernameIsValid(username string) bool { //名字长度小于20
	if len([]rune(username)) > 20 {
		return false
	}

	reg := regexp.MustCompile("[^A-Za-z0-9]") //不能有其他字符
	if reg == nil {
		fmt.Println("reg err")
	}
	allString := reg.FindAllString(username, -1)
	if allString[0] != "" {
		return false
	}

	for _, str := range util.InvalidWords { //不能包含无效词汇
		if strings.Contains(username, str) {
			return false
		}
	}
	return true
}

func PasswordIsValid(password string) bool { //密码长度小于20
	if len([]rune(password)) > 20 {
		return false
	}
	reg := regexp.MustCompile("[^A-Za-z0-9]") //不能有其他字符
	if reg == nil {
		fmt.Println("reg err")
	}
	allString := reg.FindAllString(password, -1)
	if allString[0] != "" {
		return false
	}
	return true
}
