package main

import "fmt"

func main() {
	var a int32
	var b int32

	fmt.Println("请输入两个数，以空格分隔：")
	fmt.Scanf("%d %d", &a, &b)

	if sum := a + b; sum >= 50 {
		fmt.Println("hello world!")
	}
}
