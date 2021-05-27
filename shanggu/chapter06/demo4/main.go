package main

import (
	"fmt"
	"strconv"
)

// 字符串内置函数
func main() {
	str1 := "haha"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("字符 = %c\n", str1[i])
	}

	str2 := "haha杭州"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符 = %c\n", str2[i])
	}

	// 中文占3个字节，需要转换一下
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符 = %c\n", str3[i])
	}

	// 字符串转换整数
	str4 := "123"
	num, err := strconv.Atoi(str4)
	fmt.Println(num, err)

	// 整数转字符串
	str5 := strconv.Itoa(15)
	fmt.Println(str5)

	// 字符串转byte切片
	bytes := []byte("hehe")
	fmt.Println(bytes)

	// byte切片转字符串
	str6 := string([]byte{104, 101, 104, 101})
	fmt.Println(str6)
}
