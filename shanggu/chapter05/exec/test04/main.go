package main

import "fmt"

func main() {
	// 打印1~100之间所有是9的倍数的整数的个数及总和
	num := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Printf("%d 是 9 的倍数 \n", i)
			num++
			sum += i
		}
	}
	fmt.Printf("1~100之间所有是9的倍数的整数的个数为 %d 个，总和为 %d \n", num, sum)

	k := 6
	for j := 0; j <= 6; j++ {
		if k < 0 {
			break
		}
		fmt.Printf("%d + %d = %d\n", j, k, j+k)
		k--
	}
}
