package main

import (
	"fmt"
	"time"
)

// 定时器的三种写法
func main() {
	fmt.Println("now time is", time.Now())

	// 1. time.Sleep
	time.Sleep(time.Second)
	fmt.Println("now time is", time.Now())

	// 2. time.NewTimer
	mt := time.NewTimer(time.Second)
	fmt.Println("now time is", <-mt.C)

	// 3.time.After
	fmt.Println("now time is", <-time.After(time.Second))
}
