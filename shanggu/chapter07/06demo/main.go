package main

import "fmt"

func main() {
	// 求一个数组的最大值，并得到对应的下标
	var arr = [5]int{3, 22, 1, 88, 7}
	maxIndex := 0
	maxValue := arr[0]
	for index, value := range arr {
		if value > maxValue {
			maxIndex = index
			maxValue = value
		}
	}

	fmt.Printf("数组最大值是%d, 索引是%d \n", maxValue, maxIndex)
}
