package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//创建一个新文件，写入内容 5句 "hello, Gardon"
	filepath := "D:\\code\\study-go\\shanggu\\chapter14\\demo04\\test.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err.Error())
		return
	}

	defer file.Close()

	// 先读取内容
	reader := bufio.NewReader(file)
	for {
		string, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string)
	}

	// 再写入内容
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello, Gardon 11\n")
	}

	// 刷入文件中
	writer.Flush()

	// 判断文件是否存在
	info, err := os.Stat(filepath)
	if err == nil {
		fmt.Println("文件存在：", info.Name())
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在：", info, err)
	} else {
		fmt.Println("文件未知：", info, err)
	}
}
