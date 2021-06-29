package main

import "fmt"

// 方法作用在指定类型上，不旦旦只能是结构体类型

type integer int

func (i integer) print() {
	fmt.Println("i = ", i)
}

func (i *integer) sum() {
	*i = *i + 1
}

func main() {
	var a integer = 10
	a.print()
	a.sum()
	fmt.Println(a)
}
