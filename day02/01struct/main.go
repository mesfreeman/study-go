package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName, lastName string
}

func upName(p *person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	p := person{firstName: "xiao", lastName: "he"}
	upName(&p)
	fmt.Println(p.firstName, p.lastName)

}
