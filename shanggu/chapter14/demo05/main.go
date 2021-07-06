package main

import (
	"fmt"
	"os"
)

// 命令行参数

func main() {
	fmt.Println("命令行参数有", len(os.Args))

	for idx, val := range os.Args {
		fmt.Printf("index: %v, value: %v, value type: %T\n", idx, val, val)
	}
}
