package main

import (
	"fmt"
)

// 全局定义
var (
	F1 = func() string {
		return "haha"
	}
)

// 匿名函数
func main() {
	// 方式一，直接调用
	a1 := func(a, b int) int {
		return a + b
	}(10, 11)
	fmt.Println("a1 = ", a1)

	// 方式二，赋值调用
	a2 := func(a, b int) int {
		return a - b
	}
	a22 := a2(10, 11)
	fmt.Println("a22 = ", a22)

	// 方式三，全局定义调用
	a3 := F1()
	fmt.Println("a3 = ", a3)
}
