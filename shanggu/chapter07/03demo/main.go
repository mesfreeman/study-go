package main

import "fmt"

// 演示数据的遍历
func main() {
	var person = [...]string{"宋江", "张飞", "刘备"}

	// 方式一：
	for index, value := range person {
		fmt.Printf("index is %d, value is %v \n", index, value)
	}

	fmt.Println("=======")

	// 方式二：
	for index := range person {
		fmt.Printf("index is %d, value is %v \n", index, person[index])
	}
}
