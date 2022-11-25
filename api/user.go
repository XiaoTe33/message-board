package api

import (
	"fmt"
	"message-board/dao"
)

func Register() {
	u := ""
	for {
		fmt.Print("用户名:")
		fmt.Scan(&u)
		if !UsernameIsValid(u) {
			fmt.Println("用户名过长")
			continue
		}
		if err := dao.UExist(u); err == nil {
			fmt.Println("用户名已存在")
			continue
		}
		fmt.Println("用户名合理")
		break
	}
	for {
		p := ""
		fmt.Print("密码:")
		fmt.Scan(&p)
		if !PasswordIsValid(p) {
			fmt.Println("密码过长")
			continue
		}
		dao.AddUser(u, p)
		fmt.Println("注册成功")
		return
	}
}

func Login() {
	for {
		u := ""
		p := ""
		fmt.Print("用户名:")
		fmt.Scan(&u)
		fmt.Print("密码:")
		fmt.Scan(&p)
		if err := dao.UPExist(u, p); err != nil {
			fmt.Println("用户名或密码错误")
			continue
		}
		fmt.Println("登陆成功")
		//LoginPage()
		return
	}

}
