package main

import (
	"fmt"
	"time"
)

// 有缓冲区的管道
func main() {
	ch := make(chan int, 3)

	// 起个子 go 程
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子 go 程: ", i)
		}
	}()

	// 主 go 程
	time.Sleep(time.Second * 3)
	for j := 0; j < 8; j++ {
		num := <-ch
		fmt.Println("主 go 程: ", num)
	}
}
