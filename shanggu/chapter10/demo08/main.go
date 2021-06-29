package main

import "fmt"

// 1 ) 编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
// 2 ) 声明一个方法获取立方体的体积。
// 3 ) 创建一个Box结构体变量，打印给定尺寸的立方体的体积

type Box struct {
	Length int
	Width  int
	Heigth int
}

func (b *Box) Bulk() int {
	return b.Length * b.Width * b.Heigth
}

func main() {
	var length, width, heigth int
	fmt.Println("请输入立方体的长 宽 高，依空格分隔")
	fmt.Scanf("%d %d %d", &length, &width, &heigth)

	var b Box
	b.Heigth = heigth
	b.Width = width
	b.Length = length
	fmt.Println(b.Bulk())
}
