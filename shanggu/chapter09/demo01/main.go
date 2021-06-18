package main

import "fmt"

// map基本使用
func main() {
	var m1 map[int]string
	m1 = make(map[int]string, 1)
	fmt.Println(m1)

	var m2 map[int]string = make(map[int]string, 2)
	fmt.Println(m2)

	var m3 map[int]string = map[int]string{55: "wu"}
	fmt.Println(m3)

	m4 := map[int]string{1: "haha"}
	fmt.Println(m4)
}
