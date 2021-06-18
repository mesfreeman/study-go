package main

import "fmt"

// 切片map
func main() {
	sliceMap := make([]map[string]string, 2)
	fmt.Println(sliceMap, len(sliceMap), cap(sliceMap))

	if sliceMap[0] == nil {
		sliceMap[0] = make(map[string]string, 2)
		sliceMap[0]["name"] = "xiaohe"
		sliceMap[0]["sex"] = "男"
	}
	fmt.Println(sliceMap, len(sliceMap), cap(sliceMap))

	tmpMap := map[string]string{
		"name": "zhangsan",
		"sex":  "女",
	}
	sliceMap = append(sliceMap, tmpMap)
	fmt.Println(sliceMap, len(sliceMap), cap(sliceMap))
}
