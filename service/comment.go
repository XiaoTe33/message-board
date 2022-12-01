package service

import (
	"message-board/dao"
	"message-board/model"
)

func RIDExist(rid string) bool {
	return dao.FindRID(rid)
}

func CIDExist(cid string) bool {
	return dao.FindCID(cid)
}

func FindCommentByMIDOK(mid string) map[string]interface{} {
	return dao.FindCommentByMID(mid)
}

func AddCommentOK(sender string, mid string, text string) {
	dao.AddComment(sender, mid, text)
}

func UpdateCommentByCIDOK(cid string, text string) {
	dao.UpdateCommentByCID(cid, text)
}

func DeleteCommentByCIDOK(cid string) {
	dao.DeleteCommentByCID(cid)
}

func AddResponseCommentOK(sender string, mid string, rid string, text string) {
	dao.AddResponseComment(sender, mid, rid, text)
}

func FindDialogByCIDOK(cid string) map[string]model.Comment {
	return dao.FindDialogByCID(cid)
}
