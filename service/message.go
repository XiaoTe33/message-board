package service

import (
	"message-board/util"
	"strings"
)

func TextIsValid(text string) bool { //文本长度小于80
	if len([]rune(text)) > 80 {
		return false
	}
	for _, str := range util.InvalidWords { //包含无效词汇
		if strings.Contains(text, str) {
			return false
		}
	}
	return true

}
