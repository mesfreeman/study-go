package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go setStr(ch)
	go getStr(ch)
	time.Sleep(2 * 1e9)
	fmt.Println("end")
}

func setStr(ch chan string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
}

func getStr(ch chan string) {
	for {
		str := <-ch
		fmt.Println("Get str is " + str)
	}
}
