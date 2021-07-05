package model

import "fmt"

type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 创建
func NewCustomerBy(name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 获取客户信息
func (c *Customer) GetInfo() string {
	return fmt.Sprintf("%d\t%v\t%v\t%d\t%v\t%v", c.ID, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
