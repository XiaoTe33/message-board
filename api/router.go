package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.Use(InsertCookie())
	u := r.Group("/user")
	{
		u.POST("/register", Register1)
		u.POST("/login", Login1) //登录同时显示用户的留言板
		u.POST("/changePassword", ChangePassword1)
		u.POST("/logout", Logout)
	}
	m := r.Group("/message")
	{
		m.POST("/send", Send1)
		m.POST("/change", Change1)
		m.POST("/delete", Delete1)
		m.POST("/comments", Comments)
	}
	c := r.Group("/comment")
	{
		c.POST("/send", Send2)
		c.POST("/change", Change2)
		c.POST("/delete", Delete2)
		c.POST("/response", Response)
		c.POST("/dialog", Dialog)
	}
	r.Run()
}
