package main

import "fmt"

func test01() {
	fmt.Println(11)
	defer fmt.Println("haha")
	defer fmt.Println("hehe")
	defer fmt.Println("heihei")
	fmt.Println(22)
}

func test02(x int) int {
	defer func() {
		x++
	}()
	return x
}

// defer 后进先出的逻辑
func main() {
	test01()
	fmt.Println(test02(1)) //1
}
