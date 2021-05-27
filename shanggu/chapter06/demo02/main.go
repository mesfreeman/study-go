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
	b1 := a(10)
	fmt.Println("b1 = ", b1)

	b2 := a(10)
	fmt.Println("b2 = ", b2)
}
