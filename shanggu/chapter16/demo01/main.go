package main

import (
	"fmt"
	"time"
)

// goroutine 初识

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("test() hello world")
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启一个协程

	for i := 0; i < 5; i++ {
		fmt.Println("main() hello golang")
		time.Sleep(time.Second)
	}
}
