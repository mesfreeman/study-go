package main

import "fmt"

// 冒泡排序
func BubbleSort(arr *[5]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	var arr = [...]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println(arr)
}
