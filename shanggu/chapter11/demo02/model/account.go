package model

import "fmt"

type Account struct {
	idcard   string
	password string
	blance   float32
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) SetIdcard(idcard string) {
	if len(idcard) < 6 || len(idcard) > 10 {
		fmt.Println("账号长度必须在6-10之间")
		return
	}
	a.idcard = idcard
}

func (a *Account) SetPassword(password string) {
	if len(password) != 6 {
		fmt.Println("密码必须为六位")
		return
	}
	a.password = password
}

func (a *Account) SetBlance(blance float32) {
	if blance < 20 {
		fmt.Println("余额必须>20")
		return
	}
	a.blance = blance
}
