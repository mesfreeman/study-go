package main

import "fmt"

func main() {
	// string 类型底层是byte类型的切片，但不支持重新切片赋值替换
	str := "double.he@qq.com"
	slice := str[:9]
	fmt.Println(slice)
}
