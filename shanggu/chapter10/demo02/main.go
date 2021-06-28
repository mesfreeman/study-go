package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	p1 := Person{Name: "xiaohe", Age: 10}
	fmt.Println(p1)

	var p2 *Person = new(Person)
	(*p2).Name = "xiaohong"
	p2.Age = 20 // 简化方式
	fmt.Println(*p2)
	fmt.Println(&p2.Name, &p2.Age, &p2.Sex) // 地址是连续的
}
