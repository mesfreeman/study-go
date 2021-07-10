package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 创建一个Person结构体[Name,Age,Address]
// 使用rand方法配合随机创建10个Person实例，并放入到channel中
// 遍历channel，将各个Person实例的信息显示在终端

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	rand.Seed(time.Now().Unix())
	personChan := make(chan Person, 10)

	// 随机创建
	for i := 0; i < 10; i++ {
		p := Person{
			Name:    "name_" + strconv.Itoa(i),
			Age:     rand.Intn(100),
			Address: "address_" + strconv.Itoa(i),
		}
		personChan <- p
	}

	// 遍历
	close(personChan)
	for person := range personChan {
		fmt.Println(person)
	}
}
