package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9989")
	if err != nil {
		fmt.Println("net.Dail is err", err)
		return
	}

	defer conn.Close()
	fmt.Println("客户端连接成功")

	// 接收键盘输入
	go func() {
		str := make([]byte, 1024)
		for {
			n2, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("os.Stdin.Read is err: ", err2)
				continue
			}

			// 主动写数据
			conn.Write(str[:n2])
		}
	}()

	// 接口服务器返回的数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器关闭")
			return
		}
		if err != nil {
			fmt.Printf("conn.Read is err: %v\n", err)
			return
		}
		fmt.Println("服务器回包: ", string(buf[:n]))
	}

}
