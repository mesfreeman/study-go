package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println(reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Float())
	fmt.Println(v.Interface())
	fmt.Println(v.CanSet())

	var m float64 = 3.5
	v = reflect.ValueOf(&m)
	fmt.Println(v.Type())
	v = v.Elem()
	fmt.Println(v.CanSet())
}
