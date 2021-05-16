package main

import "fmt"

func main() {
	var a float64
	var b float64

	fmt.Println("请输入两个小数，以空格分隔：")
	fmt.Scanf("%f %f", &a, &b)

	if a > 10.0 && b < 20.0 {
		fmt.Printf("%f + %f = %f\n", a, b, a+b)
	}
}
