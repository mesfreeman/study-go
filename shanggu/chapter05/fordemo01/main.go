package main

import "fmt"

func main() {
	var str string = "hello, world"

	fmt.Println("// 方式一")
	for i := 0; i < len(str); i++ {
		fmt.Printf("index is %d, value is %c \n", i, str[i])
	}

	fmt.Println("// 方式二")
	for index, value := range str {
		fmt.Printf("index is %d, value is %c \n", index, value)
	}

	fmt.Println("// 解决中英文")
	str2 := "hahaha哈哈"
	// strConvert := []rune(str2) // for range 可以自动处理中英文，但用 for 时要自己转一下之后再遍历
	for index, value := range str2 {
		fmt.Printf("index is %d, value is %c \n", index, value)
	}
}
