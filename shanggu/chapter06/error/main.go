package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil { // 说明有错误产生啦
			fmt.Println("err = ", err)
		}
	}()

	n1 := 100
	n2 := 0
	res := n1 / n2
	fmt.Println("res is : ", res)
}

func test02(num int) error {
	if num == 0 {
		return errors.New("不能为零")
	}
	return nil
}

func main() {
	test()
	err := test02(0)
	fmt.Println("err = ", err)
	fmt.Println("main()下面的代码")
}
