package main

import "fmt"

type Usb interface {
	Say()
}

type Stu struct {
}

func (s *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	// var u Usb = stu // 错误
	var u Usb = &stu
	u.Say()
	fmt.Println("here", u)

}
