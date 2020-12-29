package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("Worder start, id: %d, job: %d \n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worder end, id: %d, job: %d \n", id, j)
		result <- j * 2
	}
}

func main() {
	jobChan := make(chan int, 10)
	resultChan := make(chan int, 10)

	// 消费任务
	for i := 0; i < 10; i++ {
		go worker(i, jobChan, resultChan)
	}

	// 创建任务
	for j := 0; j < 15; j++ {
		jobChan <- j
	}
	close(jobChan)

	// 打印任务
	for k := 0; k < 15; k++ {
		<-resultChan
	}
}
