package main

import "fmt"

// 继承初探
type Student struct {
	Name  string
	Age   int
	Score int
}

// 介绍
func (s *Student) Say() {
	fmt.Println("My name is " + s.Name)
}

// 小学生
type Pupil struct {
	Student
}

// 大学生
type College struct {
	Student
}

func main() {
	pupic := &Pupil{}
	pupic.Name = "xiaohe" // 简化调用
	pupic.Say()

	college := &College{}
	college.Student.Name = "xiaohong"
	college.Say()
}
