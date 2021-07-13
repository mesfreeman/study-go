package main

import (
	"fmt"
	"reflect"
)

// 基本类型的反射
func testBaseType(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)

	fmt.Println(typ, val) // int 20

	n2 := 10
	n3 := n2 + int(val.Int())
	fmt.Println(n3)

	n4 := val.Interface().(int)
	fmt.Println(n4)
}

type Stu struct {
	Name string
	Age  int
}

// 结构体类型的反射
func testStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)

	fmt.Println(typ, val)

	s1 := val.Interface().(*Stu)
	fmt.Println(s1.Name)

	k1 := typ.Kind()
	k2 := val.Kind()
	fmt.Println(k1, k2)

	s1.Name = "zhangsan"
	fmt.Println(s1)

	val.Elem().FieldByName("Name").SetString("lisi")
	fmt.Println(s1)
}

func main() {
	var n int = 20
	testBaseType(n)

	fmt.Println("----------")

	var s1 Stu = Stu{
		Name: "xiaohe",
		Age:  10,
	}
	testStruct(&s1)
}
