package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ReadFile 读文件
func ReadFile(filePath string) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open file fail.")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputStr, err := inputReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(inputStr)
	}
}

// WriteFile 写文件
func WriteFile(srcFilePath string, toFilePath string) {
	fileByte, err := ioutil.ReadFile(srcFilePath)
	if err != nil {
		fmt.Println("Read file fail.")
		return
	}
	fmt.Println(string(fileByte))

	err = ioutil.WriteFile(toFilePath, fileByte, 0644)
	if err != nil {
		fmt.Println("Write file fail.")
		return
	}
	fmt.Println("Write file success.")
}

func main() {
	srcFilePath := `D:\code\study-go\base\22file_read_write\test.txt`
	toFilePath := `D:\code\study-go\base\22file_read_write\test_copy.txt`

	ReadFile(srcFilePath)
	WriteFile(srcFilePath, toFilePath)
}
