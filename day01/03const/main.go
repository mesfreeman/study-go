package main

import "fmt"

// 单个声明
const pi = 3.1415

const (
	a1 = iota // 0
	a2 = 10   // 10，插队赋值
	a3        // 10
	_         // 跳过赋值
	a4        // 10
)

const (
	b1 = iota // 0
	b2 = 10   // 10，插队赋值
	b3 = iota // 2
	b4        // 3
)

// 多个定义在一行时
const (
	c1, c2 = iota + 1, iota + 2 // 1, 2
	c3, c4                      // 2, 3
)

func main() {
	fmt.Println(pi)

	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)

	// fmt.Println(b1)
	// fmt.Println(b2)
	// fmt.Println(b3)
	// fmt.Println(b4)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}
