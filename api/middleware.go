package api

func TextIsValid(text string) bool { //文本长度小于80
	if len([]rune(text)) <= 80 {
		return true
	}
	return false
}

func UsernameIsValid(username string) bool { //名字长度小于20
	if len([]rune(username)) <= 20 {
		return true
	}
	return false
}

func PasswordIsValid(password string) bool { //密码长度小于20
	if len([]rune(password)) <= 20 {
		return true
	}
	return false
}
