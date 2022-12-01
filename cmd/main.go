package main

import (
	"message-board/api"
	"message-board/dao"
)

/*项目的入口*/

func main() {
	dao.InitDB()
	api.Start()
}
