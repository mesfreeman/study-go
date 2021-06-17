package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成10个整数（1-100）保存到数组，并倒序打印以及求平均值、最大值、最大值下标、并查找里面是否有55
func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for k := range arr {
		arr[k] = (rand.Intn(100) + 1)
	}
	fmt.Println(arr)

	var sum float32 = 0.0        // 总数
	var avg float32 = 0.0        // 平均值
	var maxIdx int = 0           // 最大值下标
	var maxVal int = arr[maxIdx] // 最大值
	var isHas55 bool = false     // 是否有55
	for k, v := range arr {
		if v > maxVal {
			maxIdx = k
			maxVal = v
		}
		if v == 55 {
			isHas55 = true
		}
		sum += float32(v)
	}
	avg = sum / float32(len(arr))
	fmt.Printf("总和：%v, 平均值：%v 最大值下标：%d 最大值：%v 是否有55：%v \n", sum, avg, maxIdx, maxVal, isHas55)

	// 倒序
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
