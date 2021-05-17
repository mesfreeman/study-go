package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("该整数大于零")
	} else if num < 0 {
		fmt.Println("该整数小于零")
	} else {
		fmt.Println("该整数等于零")
	}
}
