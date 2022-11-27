package service

func TextIsValid(text string) bool { //文本长度小于80
	if len([]rune(text)) <= 80 {
		return true
	}
	return false
}
