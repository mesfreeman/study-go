package main

import (
	"fmt"
	"math"
)

func main() {
	var a1 = 123.11726366
	fmt.Printf("%T\n", a1) // float64
	fmt.Printf("%f\n", a1) // 123.117264 会有一定的精度损失

	a2 := float32(13.222)  // 指定具体类型，不指定时默认为float64
	fmt.Printf("%T\n", a2) // float32
	fmt.Printf("%f\n", a2) // 13.222000

	// 最大范围
	fmt.Printf("%f\n", math.MaxFloat32)
	fmt.Printf("%f\n", math.MaxFloat64)
}
