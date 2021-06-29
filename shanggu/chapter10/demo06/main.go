package main

import "fmt"

type MethodUtils01 struct{}

type MethodUtils02 struct{}

func (m *MethodUtils01) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func (u *MethodUtils02) Print(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	var m01 MethodUtils01
	m01.Print()

	fmt.Println("-----------")
	var m02 MethodUtils02
	m02.Print(2, 3)
}
