package main

import "fmt"

// 累加器
func addTool() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 闭包
func main() {
	a := addTool()
	fmt.Println("a1 = ", a(10))
	fmt.Println("a2 = ", a(10))
	fmt.Println("a3 = ", a(10))
}
