package main

import "fmt"

func main() {
	blance := 100000.0
	t := 0
	for {
		if blance > 50000 {
			blance -= blance * 0.05
		} else {
			blance -= 1000
		}
		if blance > 0 {
			t += 1
		} else {
			break
		}
	}
	fmt.Println("总次数：", t)
}
