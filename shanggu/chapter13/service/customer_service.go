package service

import "github.com/mesfreeman/study-go/shanggu/chapter13/model"

type CustomerService struct {
	Customers   []model.Customer
	CustomerNum int
}

func (cs *CustomerService) List() []model.Customer {
	return cs.Customers
}
