package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server")

	ls, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Start server fail", err.Error())
		return
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("Accept fail", err.Error())
			return
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Conn read fail", err.Error())
			return
		}

		fmt.Println("Content is " + string(buf[:len]))
	}
}
