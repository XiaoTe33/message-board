package api

import (
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
)

func Send2(c *gin.Context) {
	sender := c.PostForm("sender")
	mid := c.PostForm("mid")
	text := c.PostForm("text")
	if !service.UsernameIsValid(sender) {
		c.JSON(200, gin.H{"err": "sender无效"})
		return
	}
	if !dao.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid无效"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "信息量过大"})
		return
	}
	dao.AddComment(sender, mid, text)
	c.JSON(200, gin.H{"msg": "评论发送成功"})
}

func Change2(c *gin.Context) {

}

func Delete2(c *gin.Context) {

}
