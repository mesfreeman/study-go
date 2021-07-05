package service

import (
	"errors"

	"github.com/mesfreeman/study-go/shanggu/chapter13/model"
)

type CustomerService struct {
	Customers   []model.Customer
	CustomerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 10, "123", "123@qq.com")
	customerService.Customers = append(customerService.Customers, *customer)
	return customerService
}

// 客户列表
func (cs *CustomerService) List() []model.Customer {
	return cs.Customers
}

// 添加客户
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.CustomerNum++
	customer.ID = cs.CustomerNum
	cs.Customers = append(cs.Customers, customer)
	return true
}

// 查询
func (cs *CustomerService) FindByID(ID int) (*model.Customer, int, error) {
	for idx, customer := range cs.Customers {
		if customer.ID == ID {
			return &customer, idx, nil
		}
	}
	return &model.Customer{}, -1, errors.New("客户不存在")
}

// 修改客户
func (cs *CustomerService) Modify(customer model.Customer, index int) bool {
	cs.Customers[index] = customer
	return true
}

// 删除客户
func (cs *CustomerService) Delete(index int) bool {
	cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
	return true
}
