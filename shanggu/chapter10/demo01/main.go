package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

// 结构体为值类型的
func main() {
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 10000

	monster2 := monster1
	monster2.Name = "蜘蛛精"

	fmt.Println("monster1: ", monster1)
	fmt.Println("monster2: ", monster2)
}
