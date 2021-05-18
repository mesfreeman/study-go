package main

import "fmt"

// 递归
func main() {
	test01(4) // 2 2 3
	test02(4) // 2
}

func test01(n int) {
	if n > 2 {
		n--
		test01(n)
	}
	fmt.Println("n01 = ", n)
}

func test02(n int) {
	if n > 2 {
		n--
		test02(n)
	} else {
		fmt.Println("n02 = ", n)
	}
}
