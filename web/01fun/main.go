package main

import "fmt"

// 函数的几种写法

func f1() {
	fmt.Println("haha")
}

func f2(x int) {
	fmt.Println(x)
}

func f3(x int, y int) {
	fmt.Println(x, y)
}

func f4(x, y int) {
	fmt.Println(x, y)
}

func f5(x, y int) int {
	return x + y
}

func f6(x, y int) (z int) {
	z = x + y
	return z
}

func f7(x, y int) (z int) {
	z = x + y
	return // 不建议这么写，生成的文档不友好
}

func f8(x, y int) (z, t int) {
	z = x
	t = y
	return z, t
}

func f9(name string, args ...int) {
	fmt.Println(name, args)
}

func main() {
	f1()
	f2(2)
	f3(3, 33)
	f4(4, 44)
	f5(5, 55)
	f6(6, 66)
	f7(7, 77)
	f8(8, 88)
	f9("haha", 1, 2, 3, 4)
}
