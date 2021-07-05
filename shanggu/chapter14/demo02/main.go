package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 缓存区读取文件内容

func main() {
	file, err := os.Open("/Users/double/projects/study-go/shanggu/chapter14/demo01/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err.Error())
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("读取结束...")
}
