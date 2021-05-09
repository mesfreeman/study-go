package main

import "fmt"

// 演示命令行参数
func main() {
	var name string
	var age int
	var sal float32
	var isPass bool

	fmt.Println("请输入你的姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入你的薪水：")
	fmt.Scanln(&sal)

	fmt.Println("请问你是否通过：")
	fmt.Scanln(&isPass)

	fmt.Printf("你输入的姓名是 %v, 年龄是 %v, 薪水是 %v, 是否通过 %v \n", name, age, sal, isPass)
}
