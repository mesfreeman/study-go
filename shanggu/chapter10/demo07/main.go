package main

import "fmt"

// 1 ) 编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float 64 类型。
// 2 ) 结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
// 3 ) 在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

type Student struct {
	Name   string
	Gender string
	Age    int
	ID     int
	Score  float64
}

func (s *Student) Say() string {
	return fmt.Sprintf("name = %s, gender = %s, age = %d, id = %d, socre = %.2f", s.Name, s.Gender, s.Age, s.ID, s.Score)
}

func main() {
	s1 := &Student{
		Name:   "xiaohe",
		Gender: "男",
		Age:    23,
		ID:     1,
		Score:  99.92,
	}
	str := s1.Say()
	fmt.Println(str)
}
