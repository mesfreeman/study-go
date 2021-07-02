package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

// 结构体切片排序，体验接口的好处

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (h HeroSlice) Len() int {
	return len(h)
}

func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	var tmp HeroSlice
	for i := 0; i < 5; i++ {
		hero := Hero{
			Name: "h" + strconv.Itoa(i),
			Age:  rand.Intn(100),
		}
		tmp = append(tmp, hero)
	}

	for _, v := range tmp {
		fmt.Println(v)
	}

	fmt.Println("===== 排序后 =====")

	sort.Sort(tmp)
	for _, v := range tmp {
		fmt.Println(v)
	}

}
