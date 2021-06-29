package main

import "fmt"

type MethodUtils01 struct{}

// 编程一个方法，方法不需要参数，在方法中打印一个 10.* 8 的矩形，在main方法中调用该方法。
func (u *MethodUtils01) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

type MethodUtils02 struct{}

// 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (u *MethodUtils02) Print(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

type MethodUtils03 struct {
	Length int
	Width  int
}

func (u *MethodUtils03) Area() int {
	return u.Length * u.Width
}

func main() {
	var m01 MethodUtils01
	m01.Print()

	fmt.Println("-----------")
	var m02 MethodUtils02
	m02.Print(2, 3)

	fmt.Println("-----------")
	m03 := &MethodUtils03{
		Length: 10,
		Width:  2,
	}
	fmt.Println("area: ", m03.Area())
}
