package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
)

func Change1(c *gin.Context) {
	mid := c.PostForm("mid")
	text := c.PostForm("text")
	if !service.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid无效"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "文本过长"})
		return
	}
	service.UpdateMessageByMIDOK(mid, text)
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
	service.AddMessageOK(sender, receiver, text)
	c.JSON(200, gin.H{"msg": "留言成功"})
	return
}

func Delete1(c *gin.Context) {
	mid := c.PostForm("mid")
	if !service.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid无效"})
		return
	}
	service.DeleteMessageByMIDOK(mid)
	c.JSON(200, gin.H{"msg": "留言删除成功"})

}

func Comments(c *gin.Context) {
	mid := c.PostForm("mid")
	if !service.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid不存在"})
		return
	}
	c.JSON(200, service.FindCommentByMIDOK(mid))
}
