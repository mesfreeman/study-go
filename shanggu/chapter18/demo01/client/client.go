package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Println("连接异常：", err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("readString err = ", err.Error())
		}

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("write err: ", err.Error())
		}
	}
}
