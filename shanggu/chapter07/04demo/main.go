package main

import "fmt"

func main() {
	// 数组的默认值
	var a1 [1]float32
	var a2 [1]int16
	var a3 [1]string
	var a4 [1]bool
	var a5 [1]byte

	fmt.Printf("a1 = %v, a2 = %v, a3 = %v, a4 = %v, a5 = %v \n", a1, a2, a3, a4, a5)
}
