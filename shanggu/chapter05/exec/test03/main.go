package main

import "fmt"

func main() {
	var year int

	fmt.Println("请输入一个年份")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400) == 0 {
		fmt.Println("该年份是闰年")
	}
}
