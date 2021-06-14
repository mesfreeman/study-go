package main

import "fmt"

func main() {
	// 定义数组的四种方式
	var a1 [3]int = [3]int{1, 2, 3}
	var a2 = [...]int{4, 5, 6}
	var a3 = [2]string{1: "a", 0: "b"}
	a4 := [3]int{7, 8}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
