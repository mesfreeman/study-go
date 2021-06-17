package main

import "fmt"

func main() {
	// 顺序查找
	arr := [...]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}

	var name string
	fmt.Println("请输入名字")
	fmt.Scanln(&name)

	for idx, val := range arr {
		if val == name {
			fmt.Printf("找到了，索引为 %d , 值为 %v ", idx, val)
			return
		}
	}

	fmt.Println("没找到~")
}
