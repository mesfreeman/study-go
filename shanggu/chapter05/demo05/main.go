package main

import "fmt"

func main() {
	var s float32
	var sex string
	fmt.Println("请输入您的百米用时及性别，以空格分隔")
	fmt.Scanf("%f %s", &s, &sex)

	// 判断
	if s < 8 {
		if sex == "男" {
			fmt.Println("恭喜您进级啦，您被分配到了男子组")
		} else {
			fmt.Println("恭喜您进级啦，您被分配到了女子组")
		}
	} else {
		fmt.Println("抱歉，您被淘汰啦~")
	}
}
