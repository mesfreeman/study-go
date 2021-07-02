package main

import "fmt"

// 演示多重继承（不推荐使用）
type Goods struct {
	Name  string
	Pirce float32
}

type Brand struct {
	Name    string
	Address string
}

// 多重继承
type Tv struct {
	Goods
	Brand
}

func main() {
	var tv Tv

	// 如果有相同属性时必须指定结构体名
	tv.Goods.Name = "电视机"
	tv.Brand.Name = "TCL"
	tv.Pirce = 2000.0 // 不同属性可以就近原则
	fmt.Println(tv)
}
