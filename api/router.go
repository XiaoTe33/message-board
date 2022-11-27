package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register1)
		u.POST("/login", Login1)
		u.POST("/changePassword", ChangePassword1)
	}
	m := r.Group("/message")
	{
		m.POST("/send", Send1)
		m.POST("/change", Change1)
		m.POST("/delete", Delete1)
	}
	c := m.Group("/comment")
	{
		c.POST("/send", Send2)
		c.POST("/change", Change2)
		c.POST("/delete", Delete2)
	}
	r.Run()
}
