package api

import (
	"github.com/gin-gonic/gin"
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
	if !service.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid无效"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "信息量过大"})
		return
	}
	service.AddCommentOK(sender, mid, text)
	c.JSON(200, gin.H{"msg": "评论发送成功"})
}

func Change2(c *gin.Context) {

	cid := c.PostForm("cid")
	text := c.PostForm("text")

	if !service.CIDExist(cid) {
		c.JSON(200, gin.H{"err": "cid不存在"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "文本过长"})
		return
	}
	service.UpdateCommentByCIDOK(cid, text)
	c.JSON(200, gin.H{"msg": "评论修改成功"})
}

func Delete2(c *gin.Context) {
	cid := c.PostForm("cid")
	if !service.CIDExist(cid) {
		c.JSON(200, gin.H{"err": "cid不存在"})
	}
	service.DeleteCommentByCIDOK(cid)
	c.JSON(200, gin.H{"msg": "评论删除成功"})
}

func Response(c *gin.Context) {
	sender := c.PostForm("sender")
	mid := c.PostForm("mid")
	rid := c.PostForm("rid")
	text := c.PostForm("text")
	if service.UsernameExist(sender) != nil {
		c.JSON(200, gin.H{"err": "sender不存在"})
		return
	}
	if !service.MIDExist(mid) {
		c.JSON(200, gin.H{"err": "mid不存在"})
		return
	}
	if !service.RIDExist(rid) {
		c.JSON(200, gin.H{"err": "rid不存在"})
		return
	}
	if !service.TextIsValid(text) {
		c.JSON(200, gin.H{"err": "文本过长"})
		return
	}
	service.AddResponseCommentOK(sender, mid, rid, text)
	c.JSON(200, gin.H{"msg": "回复成功"})
}

func Dialog(c *gin.Context) {
	cid := c.PostForm("cid")
	if !service.CIDExist(cid) {
		c.JSON(200, gin.H{"err": "cid不存在"})
		return
	}
	c.JSON(200, service.FindDialogByCIDOK(cid))
}
