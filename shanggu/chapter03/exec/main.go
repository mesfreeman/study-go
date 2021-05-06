package main

import (
	"fmt"
)

func main() {
	var a int = 10
	fmt.Printf("a address is %v \n", &a)

	var b *int = &a
	*b = 20
	fmt.Printf("a value is %v\n", a)
	fmt.Printf("b value is %v\n", *b)
	fmt.Printf("b address is %v\n", &b)
}
