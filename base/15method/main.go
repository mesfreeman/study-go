package main

import "fmt"

// Test ...
type Test struct {
	a, b int
}

func main() {
	t := new(Test)
	t.a = 12
	t.b = 13

	fmt.Println(t.a, t.b)

	t1 := t.Add()
	t2 := t.AddParam(1)
	fmt.Println(t1, t2)
}

// Add ...
func (t *Test) Add() int {
	return t.a + t.b
}

// AddParam ...
func (t *Test) AddParam(param int) int {
	return t.a + t.b + param
}
