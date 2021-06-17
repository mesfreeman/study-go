package main

import "fmt"

// 已知有个排序好（升序）的数组，插入一个元素，最后打印该数组，顺序依然为升序
func main() {
	arr1 := [6]int{22, 33, 55, 66, 77, 88}
	var arr2 [7]int
	var num int
	fmt.Println("请输入要插入的整数：")
	fmt.Scanln(&num)

	isInsert := false
	for i := 0; i < len(arr1); i++ {
		if num < arr1[i] && !isInsert {
			arr2[i] = num
			arr2[i+1] = arr1[i]
			isInsert = true
		} else if isInsert {
			arr2[i+1] = arr1[i]
		} else {
			arr2[i] = arr1[i]
		}
	}
	if !isInsert {
		arr2[6] = num
	}
	fmt.Println(arr2)
}
