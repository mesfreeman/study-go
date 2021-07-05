package main

import (
	"fmt"
	"os"
)

// 文件初识
func main() {
	file, err := os.Open("/Users/double/projects/study-go/shanggu/chapter14/demo01/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err.Error())
		return
	}

	defer file.Close()
	fmt.Println(file)
}
