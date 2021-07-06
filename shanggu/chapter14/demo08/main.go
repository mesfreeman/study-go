package main

import (
	"encoding/json"
	"fmt"
)

// json 反序列化
type Monster struct {
	Name  string  `json:"name"` // 反射机制
	Age   int     `json:"age"`
	Sal   float32 `json:"sal"`
	Skill string  `json:"skill"`
}

func main() {
	str := `{"name":"xiaohe","age":20,"sal":200000.7,"skill":"赚钱"}`
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败，err = ", err.Error())
		return
	}
	fmt.Printf("%T\t%v\n", monster, monster)
}
