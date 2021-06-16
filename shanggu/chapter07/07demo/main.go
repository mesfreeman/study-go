package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数组
	rand.Seed(time.Now().UnixNano()) // 随机种子
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println(arr)

	// 反转数组
	for i := 0; i < (len(arr) / 2); i++ {
		arr[len(arr)-1-i], arr[i] = arr[i], arr[len(arr)-1-i]
	}

	fmt.Println(arr)
}
