package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello word哈"
	byteStr := []byte(str)
	for i := 0; i < len(byteStr); i++ {
		fmt.Println(byteStr[i])
	}

	fmt.Println(utf8.RuneCountInString(str))
}
