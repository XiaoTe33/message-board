package api

import (
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
)

func Register1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !service.UsernameIsValid(username) {
		c.JSON(200, gin.H{"err": "用户名无效"})
		return
	}
	if err := dao.UsernameExist(username); err == nil {
		c.JSON(200, gin.H{"err": "用户名已存在"})
		return
	}
	if !service.PasswordIsValid(password) {
		c.JSON(200, gin.H{"err": "密码无效"})
		return
	}
	dao.AddUser(username, password)
	c.JSON(200, gin.H{"msg": "注册成功"})

}
func Login1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if err := dao.UsernameAndPasswordExist(username, password); err != nil {
		c.JSON(200, gin.H{"err": "用户名或密码错误"})
		return
	}
	c.JSON(200, dao.FindMessageByReceiver(username))
}
func ChangePassword1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if err := dao.UsernameAndPasswordExist(username, password); err != nil {
		c.JSON(200, gin.H{"err": "用户名或密码错误"})
		return
	}
	newPassword := c.PostForm("newPassword")
	if !service.PasswordIsValid(newPassword) {
		c.JSON(200, gin.H{"err": "新密码无效"})
		return
	}
	dao.UpdatePassword(username, newPassword)
	c.JSON(200, gin.H{"msg": "密码修改成功"})
}
