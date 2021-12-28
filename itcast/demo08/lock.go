package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var mutex sync.Mutex

func printer(str string) {
	mutex.Lock()
	for _, s := range str {
		fmt.Printf("%c", s)
		time.Sleep(time.Second)
	}
	mutex.Unlock()
}

func p1() {
	printer("hello")
}

func p2() {
	printer("world")
}

func main() {
	go p1()
	go p2()
	for {

	}
}
