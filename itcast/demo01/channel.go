package main

import (
	"fmt"
	"time"
)

// 无缓冲 channel
var ch = make(chan int)

// 打印机
func printer(str string) {
	for _, s := range str {
		fmt.Printf("%c", s)
		time.Sleep(time.Second)
	}
}

// person 1
func p1() {
	printer("hello")
	ch <- 9
}

// person 2
func p2() {
	<-ch
	printer("world")
}

func main() {
	go p1()
	go p2()
	for {

	}
}
