package main

import "fmt"

type CustomerView struct {
	Key  int
	Loop bool
}

func (cv *CustomerView) mainMenu() {
	for {
		fmt.Println("---------- 客户信息管理软件 ----------")
		fmt.Println("           1. 添加客户")
		fmt.Println("           2. 修改客户")
		fmt.Println("           3. 删除客户")
		fmt.Println("           4. 客户列表")
		fmt.Println("           5. 退   出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&cv.Key)
		switch cv.Key {
		case 1:
			fmt.Println("添加客户")
		case 2:
			fmt.Println("修改客户")
		case 3:
			fmt.Println("删除客户")
		case 4:
			fmt.Println("客户列表")
		case 5:
			cv.Loop = false
		default:
			fmt.Println("输入有误，请重新输入~")
		}

		if !cv.Loop {
			break
		}
	}

	fmt.Println("你退出了客户信息管理软件，欢迎下次使用~")
}

func main() {
	cv := &CustomerView{
		Key:  0,
		Loop: true,
	}

	// 显示主菜单
	cv.mainMenu()
}
