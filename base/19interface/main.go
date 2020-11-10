package main

import "fmt"

type Shper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, height float32
}

func (re Rectangle) Area() float32 {
	return re.width * re.height
}

func main() {
	sq := new(Square)
	sq.side = 5

	re := Rectangle{5, 8}

	sh := []Shper{sq, re}

	for _, o := range sh {
		fmt.Println("interface demo :", o)
		fmt.Println("interface demo :", o.Area())
	}
}
