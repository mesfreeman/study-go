package model

type student struct {
	Name string
}

// 工厂模式
func NewStudent(name string) *student {
	return &student{Name: name}
}
