package main

import "fmt"

func main() {
	var blance = 10000.0              // 余额
	var money = 0.0                   // 金额
	var note = ""                     // 说明
	var detail = "收支\t账户金额\t收支金额\t说明" // 详情

	for {
		fmt.Println("----------- 家庭收支记账软件 -----------")
		fmt.Println("             1. 收支明细")
		fmt.Println("             2. 登记收入")
		fmt.Println("             3. 登记支出")
		fmt.Println("             4. 退出软件")
		fmt.Println("请选择(1-4): ")

		var operNum int // 功能
		fmt.Scanln(&operNum)
		if operNum == 1 {
			fmt.Println("----------- 当前收支明细如下 -----------")
			if money == 0.0 {
				fmt.Println("当前没有收支明细，来一笔吧~")
				fmt.Println()
				continue
			}
			fmt.Println(detail)
		} else if operNum == 2 {
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			blance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%.2f\t%.2f\t\t%v", blance, money, note)
		} else if operNum == 3 {
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > blance {
				fmt.Println("余额不足")
				break
			}

			blance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%.2f\t%.2f\t\t%v", blance, money, note)
		} else if operNum == 4 {
			fmt.Println("退出成功，欢迎下次使用~")
			break
		} else {
			fmt.Println("请输入正确的选项")
		}

		fmt.Println()
	}
}
