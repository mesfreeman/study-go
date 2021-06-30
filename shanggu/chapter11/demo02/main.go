package main

import (
	"fmt"

	"github.com/mesfreeman/study-go/shanggu/chapter11/demo02/model"
)

func main() {
	account := model.NewAccount()
	account.SetIdcard("222333")
	account.SetPassword("112233")
	account.SetBlance(102.50)
	fmt.Println(account)
}
