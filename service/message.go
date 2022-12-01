package service

import (
	"message-board/dao"
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

func MIDExist(mid string) bool {
	return dao.FindMID(mid)
}

func DeleteMessageByMIDOK(mid string) {
	dao.DeleteMessageByMID(mid)
}

func UpdateMessageByMIDOK(mid string, text string) {
	dao.UpdateMessageByMID(mid, text)
}

func AddMessageOK(username string, receiver string, text string) {
	dao.AddMessage(username, receiver, text)
}

func FindMessageByReceiverOK(receiver string) interface{} {
	return dao.FindMessageByReceiver(receiver)
}
