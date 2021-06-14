package main

import "fmt"

func main() {
	var scores [5]float32
	for index := range scores {
		fmt.Printf("请输入第%d个成绩\n", index+1)
		fmt.Scanln(&scores[index])
	}

	for index, value := range scores {
		fmt.Printf("index is %d, value is %v\n", index, value)
	}
}
