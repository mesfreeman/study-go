package main

import "fmt"

func main() {
	var firstname, lastname string
	fmt.Println("Please input you firstname and lastname :")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi, %s %s \n", firstname, lastname)

	var input, format, i, s string
	input = "he xiao"
	format = "%s %s"
	fmt.Sscanf(input, format, &i, &s)
	fmt.Println("From the string we read: ", i, s)
}
