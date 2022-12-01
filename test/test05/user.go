package test05

import (
	"fmt"
	"message-board/dao"
	"message-board/service"
)

func Register() {
	u := ""
	for {
		fmt.Print("用户名:")
		fmt.Scan(&u)
		if !service.UsernameIsValid(u) {
			fmt.Println("用户名过长")
			continue
		}
		if err := dao.FindUsername(u); err == nil {
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
		if !service.PasswordIsValid(p) {
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
		if err := dao.FindUsernameAndPassword(u, p); err != nil {
			fmt.Println("用户名或密码错误")
			continue
		}
		fmt.Println("登陆成功")
		//LoginPage()
		return
	}
}

func ChangeP() {
	for {
		u := ""
		p := ""
		fmt.Print("用户名:")
		fmt.Scan(&u)
		fmt.Print("密码:")
		fmt.Scan(&p)
		if err := dao.FindUsernameAndPassword(u, p); err != nil {
			fmt.Println("用户名或密码错误")
			continue
		}
		fmt.Println("认证成功")
		for {
			fmt.Print("请输入新密码:")
			fmt.Scan(&p)
			if !service.PasswordIsValid(p) {
				fmt.Println("密码过长")
				continue
			}

			dao.UpdatePassword(u, p)
			break
		}
		return
	}
}
