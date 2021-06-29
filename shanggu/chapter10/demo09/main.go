package main

import "fmt"

// 1 ) 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于 18 ，收费 20 元，其它情况门票免费.
// 2 ) 请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出

type Visitor struct {
	Age int
}

func (v *Visitor) Price() int {
	if v.Age > 18 {
		return 20
	}
	return 0
}

func main() {
	v := Visitor{Age: 10}
	fmt.Println("他的票价为：", v.Price())
}
