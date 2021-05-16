package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你的年龄大于19岁了啊，你完蛋啦")
	} else {
		fmt.Println("年龄太小了啦，先放过你啦")
	}
}
