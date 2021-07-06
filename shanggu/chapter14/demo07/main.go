package main

import (
	"encoding/json"
	"fmt"
)

// json序列化
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Sal   float32 `json:"sal"`
	Skill string  `json:"skill"`
}

func main() {
	monster := Monster{
		Name:  "xiaohe",
		Age:   20,
		Sal:   200000.70,
		Skill: "赚钱",
	}

	str, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化失败，err = ", err.Error())
		return
	}
	fmt.Println(string(str))
}
