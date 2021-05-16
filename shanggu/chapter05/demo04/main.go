package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("请输入三个小数，以空格分隔")
	fmt.Scanf("%f %f %f", &a, &b, &c)

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		x2 := (-b - math.Sqrt(b*b-4*a*c)) / 2 * a
		fmt.Println("该方程有两个解，分别是：", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		fmt.Println("该方程只有一个解，值是：", x1)
	} else {
		fmt.Println("该方程没有解")
	}
}
