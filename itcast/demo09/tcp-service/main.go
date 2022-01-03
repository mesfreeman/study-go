package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9989")
	if err != nil {
		fmt.Println("net.Listen is err, ", err)
		return
	}

	fmt.Println("服务器等待客户端建立连接...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept is err, ", err)
			return
		}
		fmt.Println("服务器与客户端建立连接成功...")

		// 开启协程
		go handlerConnect(conn)
	}
}

// 处理服务器与客户端的数据通信
func handlerConnect(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println(addr.String(), "客户端连接成功 ")

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 || string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
			fmt.Println(addr, "客户端断开连接")
			return
		}

		if err != nil {
			fmt.Println("conn.Read is err", err)
			return
		}

		// 回发给客户端
		fmt.Println("服务器读到数据", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
