package main

import "fmt"

// 单个声明
// var name string
// var age int
// var isPeople bool

// 批量声明
var (
	name     string
	age      int
	isPeople bool
)

func main() {
	name = "小明"
	age = 10
	isPeople = true
	sex := 1
	var height int = 160

	fmt.Println(name, age, isPeople)
	fmt.Printf("sex is %d \n", sex)
	fmt.Printf("height is %d \n", height)
}
