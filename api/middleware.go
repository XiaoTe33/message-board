package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		if c.FullPath() == "/user/login" || c.FullPath() == "/user/register" {
			c.Next()
		} else {
			fmt.Println(c.FullPath())
			cookie, err := c.Cookie("cookie")
			if err != nil {
				fmt.Println("cookie err")
				cookie = "NotSet"
				c.Abort()
				c.JSON(200, gin.H{"err": "请先登录"})
			} else {
				if c.FullPath() == "/user/logout" {

				} else {
					if cookie != username {
						c.JSON(200, gin.H{"err": "请先登录"})
						c.Abort()
					}

					if cookie == username {
						c.JSON(200, gin.H{"msg": "当前用户【" + username + "】"})
						c.Next()
					}
				}
			}
		}
	}
}
