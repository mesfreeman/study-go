package main

import "fmt"

func feibo(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			return
		}
	}
}

// 斐波那契数列
func main() {
	ch := make(chan int, 3)
	quit := make(chan bool)

	go feibo(ch, quit)

	x, y := 1, 1
	for i := 1; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
