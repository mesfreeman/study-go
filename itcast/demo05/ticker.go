package main

import (
	"fmt"
	"time"
)

// 周期定时
func main() {
	fmt.Println("now time is ", time.Now())

	// 创建周期定时器
	myTicker := time.NewTicker(time.Second)

	go func() {
		for {
			nowTime := <-myTicker.C
			fmt.Println("now time is", nowTime)
		}
	}()

	for {
	}
}
