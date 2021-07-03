package main

import "fmt"

// 类型断言

type Point struct {
	X int
	Y int
}

func main() {
	var a interface{}
	var point = Point{
		X: 20,
		Y: 30,
	}

	a = point
	fmt.Println(a)

	var b Point
	// b = a // 不可以会报错
	b = a.(Point) // 断言
	fmt.Println(b)

	// 类型断言如果不匹配会报错
	var c interface{}
	var d float32
	c = d
	var e float64
	// e = c.(float32) // 报错
	fmt.Println(c, e)
}
