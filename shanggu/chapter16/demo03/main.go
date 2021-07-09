package main

import (
	"fmt"
	"sync"
	"time"
)

// 利用协程计算指定200内数的阶乘

var (
	tmp  = make(map[int]int64, 10)
	lock sync.Mutex
)

// 阶乘
func test(num int) {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	lock.Lock() // 加锁：fatal error: concurrent map iteration and map write
	tmp[num] = int64(result)
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// @todo 此方法不合理
	time.Sleep(time.Second * 5)

	for key, value := range tmp {
		fmt.Printf("key = %v, value = %v \n", key, value)
	}
}
