package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听")

	ln, err := net.Listen("tcp", ":8989")
	if err != nil {
		fmt.Println("监听异常：", err.Error())
		return
	}
	defer ln.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("接收异常：", err.Error())
		}

		fmt.Println("conn come in, ip: ", conn.RemoteAddr())

		go proccess(conn)
	}
}

func proccess(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出：", err.Error())
			return
		}
		fmt.Print(conn.RemoteAddr().String() + ": " + string(buf[:len]))
	}
}
