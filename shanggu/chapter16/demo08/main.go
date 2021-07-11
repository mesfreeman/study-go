package main

import (
	"fmt"
	"strconv"
)

// select 在协程里的使用

func main() {
	intChan := make(chan int, 5)
	stringChan := make(chan string, 8)

	for i := 0; i < 5; i++ {
		intChan <- i
	}

	for i := 0; i < 8; i++ {
		stringChan <- "hello " + strconv.Itoa(i)
	}

label: // 跳出标记
	for {
		select {
		case v := <-intChan:
			fmt.Println("intCha ", v)
		case v := <-stringChan:
			fmt.Println("stringChan ", v)
		default:
			fmt.Println("都取不到值")
			break label
		}
	}

	fmt.Println("end")
}
