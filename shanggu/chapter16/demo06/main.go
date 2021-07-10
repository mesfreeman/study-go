package main

import (
	"fmt"
)

// 开启一个writeData协程，向管道intChan中写入50个整数
// 开启一个readData协程，从管道intChan中读取writeData写入的数据
// 注意：writeData和readData操作的是同一个管道
// 主线程需要等待writeData和readData协程都完成工作才能退出管道

func writeNum(intChan chan int) {
	for i := 1; i < 50; i++ {
		intChan <- i
		fmt.Println("写入：", i)
	}
	close(intChan)
}

func readNum(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取：", val)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)

	go writeNum(intChan)
	go readNum(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println("读写结束")
}
