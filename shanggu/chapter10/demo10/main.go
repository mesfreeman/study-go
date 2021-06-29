package main

import (
	"fmt"

	"github.com/mesfreeman/study-go/shanggu/chapter10/demo10/model"
)

// 工厂模式的使用
func main() {
	s := model.NewStudent("xiaohe")
	fmt.Println(s.Name)
}
