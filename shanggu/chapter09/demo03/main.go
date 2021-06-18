package main

import "fmt"

// map 的 增 删 改 查 遍历
func main() {
	students := map[string]string{
		"s1": "zhangsan",
		"s2": "lisi",
	}

	// 增
	students["s3"] = "wangwu"
	fmt.Println(students)

	// 删
	delete(students, "s3")
	fmt.Println(students)

	// 改
	students["s2"] = "haha"
	fmt.Println(students)

	// 查
	val, ok := students["s1"]
	if ok {
		fmt.Println("找到了，值为：", val)
	} else {
		fmt.Println("没找到")
	}

	// 遍历
	for key, val := range students {
		fmt.Println("key: ", key, "val: ", val)
	}
}
