package main

import "fmt"

// chanenl 的基本使用

type Cat struct {
	Name string
	Age  int
}

func main() {
	// int类的的管道
	var intChan chan int

	// 使用前必须初始化
	intChan = make(chan int, 3)

	// 添加
	intChan <- 10
	intChan <- 20
	intChan <- 30
	// intChan <- 40 // 容量为3时只能放3个元素，否则会报错

	// 打印长度及容量
	fmt.Printf("inChan len is %v, cap is %v\n", len(intChan), cap(intChan))

	var num1 int
	num1 = <-intChan // 取元素，先进先出的逻辑
	fmt.Println("num1 = ", num1)

	// map类型管理
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := map[string]string{
		"name":   "xiaohe",
		"gender": "男",
	}
	m2 := map[string]string{
		"name":   "zhangsan",
		"gender": "男",
	}
	mapChan <- m1
	mapChan <- m2
	fmt.Println(<-mapChan)

	// 结构体类型
	var catChat chan Cat
	catChat = make(chan Cat, 5)
	cat1 := Cat{
		Name: "haha",
		Age:  10,
	}
	cat2 := Cat{
		Name: "hehe",
		Age:  20,
	}
	catChat <- cat1
	catChat <- cat2
	fmt.Println(<-catChat)

}
