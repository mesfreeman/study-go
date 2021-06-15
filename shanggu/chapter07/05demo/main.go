package main

import "fmt"

func test(arr [2]int) {
	arr[0] = 88
}

// 引用传递
func test02(arr *[2]int) {
	arr[1] = 77
}

func main() {
	// 数组是值传递
	arr := [2]int{1, 2}
	test(arr)
	test02(&arr)
	fmt.Println(arr)
}
