package main

import (
	"fmt"
	"time"
)

func addNum(numChan chan<- int) {
	for i := 1; i <= 200; i++ {
		numChan <- i
	}
	close(numChan)
}

func readNum(numChan <-chan int, resChan chan<- map[int]int, exitChan chan<- bool) {
	for {
		var res int
		num, ok := <-numChan
		if !ok {
			exitChan <- true
			break
		}

		for i := 1; i <= num; i++ {
			res += i
		}

		resChan <- map[int]int{num: res}
	}
}

func main() {
	numChan := make(chan int, 500)
	go addNum(numChan)

	resChan := make(chan map[int]int, 200)
	exitChan := make(chan bool, 8)

	starTime := time.Now()
	for i := 0; i < 8; i++ {
		go readNum(numChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
		close(exitChan)
		fmt.Println("耗时：", time.Since(starTime).Nanoseconds())
	}()

	for val := range resChan {
		fmt.Println(val)
	}
}
