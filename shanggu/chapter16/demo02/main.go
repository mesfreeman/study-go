package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpu = ", cpu)

	// 设置cup数
	runtime.GOMAXPROCS(cpu - 1)
	fmt.Println("ok")
}
