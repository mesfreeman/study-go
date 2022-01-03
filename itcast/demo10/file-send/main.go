package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 文件发送端
func main() {
	// 输入要传输的文件
	list := os.Args
	if len(list) != 2 {
		fmt.Println("请输入文件，格式为: go run xxx.go filename")
		return
	}

	// 获取文件名称
	fileInfo, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("os.Stat is err, ", err)
		return
	}

	filename := fileInfo.Name()

	// 发送连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8876")
	if err != nil {
		fmt.Println("net.Dial is err, ", err)
		return
	}
	defer conn.Close()

	// 发送文件名到接收端
	_, err = conn.Write([]byte(filename))
	if err != nil {
		fmt.Println("conn.Write is err, ", err)
		return
	}

	// 接收接收端的回包数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read is err, ", err)
		return
	}

	// 文件名发送成功，开始发送文件内容
	if string(buf[:n]) == "ok" {
		sendFile(conn, list[1])
	}
}

// 发送文件内容
func sendFile(conn net.Conn, filepath string) {
	// 打开文件
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("os.Open is err: ", err)
		return
	}
	defer f.Close()

	// 读取文件内容并写到接收端，读多少写多少
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送完成")
			} else {
				fmt.Println("f.Read is err: ", err)
			}
			return
		}

		// 写到网络中
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write is err: ", err)
			return
		}
	}
}
