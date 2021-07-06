package main

import (
	"flag"
	"fmt"
)

// 命令行参数 flag 包的使用

func main() {
	var (
		user     string
		password int
	)

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.IntVar(&password, "p", 123, "密码，默认为123")
	flag.Parse() // 一定要写这个操作

	fmt.Println(user, password)
}
