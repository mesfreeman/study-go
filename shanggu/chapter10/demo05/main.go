package main

import "fmt"

type Person struct {
	Name string
}

// 如果一个类型实现了String()这个方法，则打印时会调用此方法
func (p *Person) String() string {
	return fmt.Sprintf("name = %s", p.Name)
}

func main() {
	p := Person{
		Name: "xiaohe",
	}
	fmt.Println(&p)
}
