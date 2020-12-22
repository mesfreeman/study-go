package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	fmt.Println("This is f1")
}

func f2() {
	defer wg.Done()
	fmt.Println("This is f2")
}

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go f1()
		go f2()
	}
	wg.Wait()
}
