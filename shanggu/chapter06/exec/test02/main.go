package main

import "fmt"

/*
    *
   ***
  *****
*/

func jinZiTa(raw int) {
	for i := 1; i <= raw; i++ {
		// 打印空格
		for j := 1; j <= raw-i; j++ {
			fmt.Print(" ")
		}

		// 打印星号
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var raw int
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&raw)
	jinZiTa(raw)
}
