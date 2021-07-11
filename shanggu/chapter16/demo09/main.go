package main

import (
	"fmt"
	"time"
)

// 协程中的异常捕获处理机制

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() ", i)
		time.Sleep(time.Second)
	}
}

func test2() {
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test2() 出错了")
		}
	}()

	var m1 map[int]string

	// 这个会出错，因为没有make空间，
	// panic: assignment to entry in nil map
	m1[0] = "xiaohe"
}

func main() {
	go test1()
	go test2()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ", i)
		time.Sleep(time.Second)
	}
}
