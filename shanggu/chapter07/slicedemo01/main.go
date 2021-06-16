package main

import "fmt"

func main() {
	// 切片的基本使用
	var arr1 = [4]int{4, 22, 7, 8}

	// 切片创建方式一：
	arr2 := arr1[1:2]

	fmt.Println("arr1 ", arr1)
	fmt.Println("arr1 len ", len(arr1))

	fmt.Println("arr2", arr2)
	fmt.Println("arr2 len ", len(arr2))
	fmt.Println("arr2 cap ", cap(arr2))

	// 切片创建方式二：
	arr3 := make([]int, 3, 4)
	fmt.Println("arr3", arr3, " len ", len(arr3), " cap ", cap(arr3))
	arr3[0] = 1
	arr3[1] = 33
	fmt.Println("arr3", arr3, " len ", len(arr3), " cap ", cap(arr3))

	// 切片创建方式三：
	arr4 := []int{6, 88, 2}
	fmt.Println("arr4", arr4, " len ", len(arr4), " cap ", cap(arr4))
}
