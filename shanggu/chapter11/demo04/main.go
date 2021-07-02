package main

import "fmt"

// 匿名结构为基础数据类型时
type Stu struct {
	Name string
	int
}

func main() {
	var stu Stu
	stu.Name = "xiaohe"
	stu.int = 123 // 直接调用基础数据类型即可，但只能存在一个，否则必须指定名字
	fmt.Println(stu)
}
