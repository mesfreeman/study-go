package main

import "fmt"

func main() {
	var username, password string
	for times := 3; times > 0; times-- {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)

		fmt.Println("请输入密码：")
		fmt.Scanln(&password)

		if username == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		}
		if times-1 > 0 {
			fmt.Printf("登录失败，你还有 %d 次机会\n", times-1)
		}
	}

	fmt.Println("登录失败")
}
