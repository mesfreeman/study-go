package main

import "fmt"

func main() {
	// 字符串
	a1 := "hello"
	fmt.Println(a1)

	// 字符
	a2 := 'A'
	fmt.Println(a2)

	// 反引号原样输出
	a3 := `
		云想衣裳花想容
		春风扶柳露华容
	`
	fmt.Println(a3)

	// 拼接
	a4 := "梦想还是要有的"
	a5 := "万一实现了呢"
	fmt.Println(a4 + "，" + a5)
	fmt.Println(len(a4))

	a6 := fmt.Sprintf("%s%s%s", a4, "，", a5) // 拼接并返回变量
	fmt.Println(a6)

	a7 := 3.14
	a8 := int(a7)
	fmt.Println(a8)
}
