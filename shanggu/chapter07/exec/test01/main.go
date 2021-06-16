package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 4, 44, 22}
	slice2 := []int{88}
	fmt.Println("slice1 = ", slice1, "slice2 = ", slice2)
	num := copy(slice1, slice2)
	fmt.Println("slice1 = ", slice1, "slice2 = ", slice2, "num = ", num)
}
