package main

import "fmt"

func main() {
	age := 12
	if age > 60 {
		fmt.Println("老年人")
	} else if age >= 18 {
		fmt.Println("青壮年")
	} else {
		fmt.Println("未成年")
	}

	if age := 30; age > 60 {
		fmt.Println("老年人")
	} else if age >= 18 {
		fmt.Println("青壮年")
	} else {
		fmt.Println("未成年")
	}
}
