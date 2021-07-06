package main

import (
	"fmt"
	"io/ioutil"
)

// 文件一次性读取

func main() {
	content, err := ioutil.ReadFile("D:\\code\\study-go\\shanggu\\chapter14\\demo03\\test.txt")
	if err != nil {
		fmt.Println("read file err = ", err.Error())
		return
	}

	// 内部封装了文件的打开与关闭操作，因此不需要处理
	// 注意：只适合文件比较小的时候

	fmt.Println(string(content))
}
