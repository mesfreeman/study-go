package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 结构体序列化
func main() {
	p1 := Person{
		Name: "xiaohe",
		Age:  20,
	}

	fmt.Println(p1)

	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}

	fmt.Println(string(bytes))

}
