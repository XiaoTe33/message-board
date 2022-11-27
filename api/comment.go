package api

import (
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
)

func Change1(c *gin.Context) {
	mid := c.PostForm("mid")
	text := c.PostForm("text")
	if !dao.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid无效"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "文本过长"})
		return
	}
	dao.UpdateMessageByMID(mid, text)
	c.JSON(200, gin.H{"msg": "修改成功"})
	return
}

func Send1(c *gin.Context) {
	sender := c.PostForm("sender")
	receiver := c.PostForm("receiver")
	text := c.PostForm("text")
	if !service.UsernameIsValid(sender) {
		c.JSON(200, gin.H{"err": "sender无效"})
		return
	}
	if !service.UsernameIsValid(receiver) {
		c.JSON(200, gin.H{"err": "receiver无效"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "信息量过大"})
		return
	}
	dao.AddMessage(sender, receiver, text)
	c.JSON(200, gin.H{"msg": "留言成功"})
	return
}

func Delete1(c *gin.Context) {
	mid := c.PostForm("mid")
	if !dao.MIDExist(mid) {
		c.JSON(200, gin.H{"msg": "mid无效"})
		return
	}
	dao.DeleteMessageByMID(mid)

}
