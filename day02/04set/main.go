package main

import "fmt"

type Person struct {
	firstName, lastName string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func main() {
	p := new(Person)
	// p.firstName = "xiao"
	p.SetFirstName("xiao")

	fmt.Println(p.FirstName())
}
