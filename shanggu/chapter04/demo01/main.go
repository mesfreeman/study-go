package main

import "fmt"

func main() {
	// 相除时，如果为整数，则结果也为整数
	fmt.Println("10 / 3", 10/3) // 3

	// 如果想结果保留小数，则需求有浮点数参数运算
	fmt.Println("10.0 / 3 =", 10.0/3) // 3.225806451612903
	fmt.Println("10 / 3.1 =", 10/3.1) // 3.225806451612903

	// 演示 % 使用
	fmt.Println("10 % 3 =", 10%3)     // 1
	fmt.Println("-10 % 3 =", -10%3)   // -1
	fmt.Println("10 % -3 =", 10%-3)   // 1
	fmt.Println("-10 % -3 =", -10%-3) // -1

	day := 97
	fmt.Println("还有", day/7, "个星期零", day%7, "天放假")
}
