package main

import (
	"fmt"
	"net"
	"os"
)

// 文件接收端
func main() {
	// 建立网络连接
	listener, err := net.Listen("tcp", "127.0.0.1:8876")
	if err != nil {
		fmt.Println("net.Listen is err: ", err)
		return
	}

	// 阻塞监听文件传输
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept is err: ", err)
		return
	}
	defer conn.Close()

	// 接收文件名称
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read is err: ", err)
		return
	}
	filename := string(buf[:n])

	// 回包 OK
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn.Write is err: ", err)
		return
	}

	// 接收文件
	receiveFile(conn, filename)
}

// 接收文件
func receiveFile(conn net.Conn, filename string) {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create is err: ", err)
		return
	}

	// 接收文件内容并写入文件中
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成")
			return
		}

		// 写入文件中
		_, err = file.Write(buf[:n])
		fmt.Println("file.Write is err: ", err)
	}
}
