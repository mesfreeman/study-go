package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	for i := 0; i < 8; i++ {
		fmt.Println("生产者:", i*i)
		ch <- i * i
	}
	close(ch) // 关闭
}

// 消费者
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("消费者:", num)
		time.Sleep(time.Second)
	}
}

// 生产者、消费者模型示例
func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}
