package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
)

func Register1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !service.UsernameIsValid(username) {
		c.JSON(200, gin.H{"err": "用户名无效"})
		return
	}
	if err := service.UsernameExist(username); err == nil {
		c.JSON(200, gin.H{"err": "用户名已存在"})
		return
	}
	if !service.PasswordIsValid(password) {
		c.JSON(200, gin.H{"err": "密码无效"})
		return
	}
	service.AddUserOK(username, password)
	c.JSON(200, gin.H{"msg": "注册成功"})

}

func Login1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if err := service.UsernameAndPassWordExist(username, password); err != nil {
		c.JSON(200, gin.H{"err": "用户名或密码错误"})
		return
	}
	c.SetCookie("cookie", username, 3600, "/user/", "localhost", false, true)
	c.SetCookie("cookie", username, 3600, "/message/", "localhost", false, true)
	c.SetCookie("cookie", username, 3600, "/comment/", "localhost", false, true)
	c.JSON(200, service.FindMessageByReceiverOK(username))

}

func ChangePassword1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if err := service.UsernameAndPassWordExist(username, password); err != nil {
		c.JSON(200, gin.H{"err": "用户名或密码错误"})
		return
	}
	newPassword := c.PostForm("newPassword")
	if !service.PasswordIsValid(newPassword) {
		c.JSON(200, gin.H{"err": "新密码无效"})
		return
	}
	service.UpdatePasswordOK(username, newPassword)
	c.JSON(200, gin.H{"msg": "密码修改成功"})
}

func Logout(c *gin.Context) {
	cookie, err := c.Cookie("cookie")
	if err != nil {
		fmt.Println("Cookie err")
	}
	c.SetCookie("cookie", "@我是注销用的", -1, "/user/", "localhost", false, true)
	c.SetCookie("cookie", "@我是注销用的", -1, "/comment/", "localhost", false, true)
	c.SetCookie("cookie", "@我是注销用的", -1, "/message/", "localhost", false, true)
	c.JSON(200, gin.H{"msg": "【" + cookie + "】账号已注销"})
}
