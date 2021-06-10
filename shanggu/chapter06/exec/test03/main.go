package main

import (
	"fmt"
)

// 循环输入年月日，判断日期是否正确
func main() {
	for {
		var (
			year  int
			month int
			day   int
		)
		fmt.Println("请输入年份：")
		fmt.Scanln(&year)
		fmt.Println("请输入月份：")
		fmt.Scanln(&month)
		if month < 0 || month > 12 {
			fmt.Println("月份不正确！")
			continue
		}
		fmt.Println("请输入日期：")
		fmt.Scanln(&day)
		fmt.Printf("您输入的日期是，%v年%v月%v日\n", year, month, day)
	}
}
