package main

import "fmt"

// 水仙花数
func main() {
	var number int
	fmt.Println("请输入一个三位整数：")
	fmt.Scanln(&number)

	// 153 = 1*1*1 + 5*5*5 + 3*3*3
	var a, b, c int
	a = number / 100
	b = number / 10 % 10
	c = number % 10

	if number == a*a*a+b*b*b+c*c*c {
		fmt.Println("该整数是水仙花数")
	} else {
		fmt.Println("该整数不是水仙花数")
	}
}
