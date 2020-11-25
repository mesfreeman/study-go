package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")

	go longWait()
	go shortWait()

	fmt.Println("Start main sleep")
	time.Sleep(8 * 1e9)
	fmt.Println("End main sleep")
}

func longWait() {
	fmt.Println("Start long wait")
	time.Sleep(5 * 1e9)
	fmt.Println("End long wait")
}

func shortWait() {
	fmt.Println("Start short wait")
	time.Sleep(2 * 1e9)
	fmt.Println("End short wait")
}
