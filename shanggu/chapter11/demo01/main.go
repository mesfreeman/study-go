package main

import "fmt"

// 面向对象的抽象演示：银行账户相关操作
type Account struct {
	Idcard   string
	Password string
	Blance   float32
}

// 存款
func (a *Account) Deposed(password string, money float32) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}

	if money <= 0.0 {
		fmt.Println("金额必须大于等于零")
		return
	}

	a.Blance += money
	fmt.Printf("存款成功，账户：%v 当前金额：%v \n", a.Idcard, a.Blance)
}

// 取款
func (a *Account) Withdraw(password string, money float32) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}

	if money <= 0.0 {
		fmt.Println("金额必须大于等于零")
		return
	}
	if money > a.Blance {
		fmt.Println("余额不足")
		return
	}

	a.Blance -= money
	fmt.Printf("取款成功，账户：%v 当前金额：%v \n", a.Idcard, a.Blance)
}

// 取款
func (a *Account) Info(password string) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}

	fmt.Printf("账户：%v 当前余额：%v \n", a.Idcard, a.Blance)
}

func main() {
	account := &Account{
		Idcard:   "88888888",
		Password: "123",
		Blance:   20.30,
	}

	for {
		operNum := 0
		password := ""
		money := 0.0
		fmt.Println("请输入你要做的操作：1-查询 2-存款 3-取款 9-退出")
		fmt.Scanln(&operNum)

		// 操作控制
		if operNum == 1 {
			fmt.Println("请输入你的密码：")
			fmt.Scanln(&password)
			account.Info(password)
			continue
		} else if operNum == 2 {
			fmt.Println("请输入你的密码：")
			fmt.Scanln(&password)
			fmt.Println("请输入存款金额：")
			fmt.Scanln(&money)
			account.Deposed(password, float32(money))
			continue
		} else if operNum == 3 {
			fmt.Println("请输入你的密码：")
			fmt.Scanln(&password)
			fmt.Println("请输入取款金额：")
			fmt.Scanln(&money)
			account.Withdraw(password, float32(money))
			continue
		} else if operNum == 9 {
			fmt.Println("本次服务结束，谢谢光临~")
			break
		} else {
			fmt.Println("指令错误，请重新输入")
			continue
		}
	}
}
