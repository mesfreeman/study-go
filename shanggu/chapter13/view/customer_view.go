package main

import (
	"fmt"

	"github.com/mesfreeman/study-go/shanggu/chapter13/model"
	"github.com/mesfreeman/study-go/shanggu/chapter13/service"
)

type CustomerView struct {
	Key             int
	Loop            bool
	customerService *service.CustomerService
}

// 显示客户列表
func (cv *CustomerView) list() {
	fmt.Println("---------- 客户列表 ----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	// 遍历列表数据
	customers := cv.customerService.List()
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
	}

	fmt.Printf("\n ---------- 加载完成 ----------\n\n")
}

// 添加客户
func (cv *CustomerView) add() {
	fmt.Println("---------- 添加客户 ----------")
	var (
		name   string
		gender string
		age    int
		phone  string
		email  string
	)

	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	fmt.Scanln(&email)

	customer := model.NewCustomerBy(name, gender, age, phone, email)
	if cv.customerService.Add(*customer) {
		fmt.Println("---------- 添加完成 ----------")
	} else {
		fmt.Println("---------- 添加失败 ----------")
	}
}

// 修改客户
func (cv *CustomerView) modify() {
	fmt.Println("---------- 修改客户 ----------")
	for {
		fmt.Print("请选择待修改的客户编号(-1退出)：")
		var modifyId int
		fmt.Scanln(&modifyId)
		if modifyId == -1 {
			fmt.Println("退出操作")
			break
		}

		customer, idx, err := cv.customerService.FindByID(modifyId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var (
			newName   string
			newGender string
			newAge    int
			newPhone  string
			newEmail  string
		)

		fmt.Printf("姓名(%v)：", customer.Name)
		fmt.Scanln(&newName)
		fmt.Printf("性别(%v)：", customer.Gender)
		fmt.Scanln(&newGender)
		fmt.Printf("年龄(%v)：", customer.Age)
		fmt.Scanln(&newAge)
		fmt.Printf("电话(%v)：", customer.Phone)
		fmt.Scanln(&newPhone)
		fmt.Printf("邮件(%v)：", customer.Email)
		fmt.Scanln(&newEmail)

		if newName != "" {
			customer.Name = newName
		}
		if newGender != "" {
			customer.Gender = newGender
		}
		if newAge != 0 {
			customer.Age = newAge
		}
		if newPhone != "" {
			customer.Phone = newPhone
		}
		if newEmail != "" {
			customer.Email = newEmail
		}

		if cv.customerService.Modify(*customer, idx) {
			fmt.Println("---------- 修改成功 ----------")
		} else {
			fmt.Println("---------- 修改失败 ----------")
		}
	}
}

// 删除客户
func (cv *CustomerView) delete() {
	fmt.Println("---------- 删除客户 ----------")
	var deleteId int
	for {
		fmt.Print("请选择待修改的客户编号(-1退出)：")
		fmt.Scanln(&deleteId)
		if deleteId == -1 {
			fmt.Println("退出操作")
			break
		}

		_, idx, err := cv.customerService.FindByID(deleteId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if cv.customerService.Delete(idx) {
			fmt.Println("---------- 删除成功 ----------")
		} else {
			fmt.Println("---------- 删除失败 ----------")
		}
	}
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
			cv.add()
		case 2:
			cv.modify()
		case 3:
			cv.delete()
		case 4:
			cv.list()
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

	// 初始化
	cv.customerService = service.NewCustomerService()

	// 显示主菜单
	cv.mainMenu()
}
