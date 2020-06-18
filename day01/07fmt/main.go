package main

import "fmt"

func main() {
	a1 := 100
	fmt.Printf("%T\n", a1) // 数据类型
	fmt.Printf("%v\n", a1) // 数据值 100
	fmt.Printf("%b\n", a1) // 二进制
	fmt.Printf("%d\n", a1) // 十进制
	fmt.Printf("%o\n", a1) // 八进制
	fmt.Printf("%x\n", a1) // 十六进制

	a2 := "hello"
	fmt.Printf("%s\n", a2)  // 字符串值 hello
	fmt.Printf("%#v\n", a2) // 数据值 "hello"
	fmt.Printf("%c\n", 'B') // 数据值 B
}
