package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // 十进制 101
	fmt.Printf("%b\n", i1) // 十进制转二进制 1100101
	fmt.Printf("%o\n", i1) // 十进制转八进制 145
	fmt.Printf("%x\n", i1) // 十进制转十六进制 65

	// 八进制，以0开头
	i2 := 066
	fmt.Printf("%d\n", i2) // 十进制 54

	// 十六进制，以0x开头
	i3 := 0x111
	fmt.Printf("%d\n", i3) // 十进制 273

	// 查看变量类型
	fmt.Printf("%T\n", i1) // int，在不指定时默认为int类型

	// 声明int8类型
	var i4 int8 = 12
	fmt.Printf("%d\n", i4) // 12
	fmt.Printf("%T\n", i4) // int8

	// 第二种声明方式
	i5 := int8(13)
	fmt.Printf("%d\n", i5) // 13
	fmt.Printf("%T\n", i5) // int13
}
