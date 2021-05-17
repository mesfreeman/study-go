package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("该数为：", i, sum)
			break
		}
	}
}
