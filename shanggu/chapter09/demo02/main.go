package main

import "fmt"

/*
课堂练习：演示一个key-value 的value是map的案例
比如：我们要存放3个学生信息, 每个学生有 name和sex 信息
*/
func main() {
	// 方式一：
	var students1 = make(map[string]map[string]string, 3)
	students1["s1"] = make(map[string]string, 2)
	students1["s1"]["name"] = "xiaohe"
	students1["s1"]["sex"] = "男"
	students1["s2"] = make(map[string]string, 2)
	students1["s2"]["name"] = "zhangsan"
	students1["s2"]["sex"] = "男"
	students1["s3"] = make(map[string]string, 2)
	students1["s3"]["name"] = "xiaohong"
	students1["s3"]["sex"] = "女"
	fmt.Println("students1", students1)

	// 方式二：
	students2 := map[string]map[string]string{
		"s1": {
			"name": "xiaohe",
			"sex":  "男",
		},
		"s2": {
			"name": "zhangsan",
			"sex":  "男",
		},
		"s3": {
			"name": "xiaohong",
			"sex":  "女",
		},
	}
	fmt.Println("students2", students2)
}
