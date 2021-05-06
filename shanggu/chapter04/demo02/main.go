package main

import "fmt"

func test() int {
	return 2
}

func main() {
	// 基础赋值
	var a int
	a = 10
	fmt.Println("a = ", a)

	// 变量赋值
	var b int = 20
	var c int = 30
	fmt.Println("交换前 b = ", b, " c = ", c)

	var d int
	d = b
	b = c
	c = d
	fmt.Println("交换后 b = ", b, " c = ", c)

	// 复合赋值
	c += 7
	fmt.Println("c = ", c)

	// 表达式
	c = 2 + 5*7
	fmt.Println("c = ", c)

	// 函数
	c = 2 + test()
	fmt.Println("c = ", c)

	// 不使用第三个变量进行交换
	dd := 22
	cc := 33
	dd, cc = cc, dd
	fmt.Println("交换后 dd = ", dd, " cc = ", cc)
}
