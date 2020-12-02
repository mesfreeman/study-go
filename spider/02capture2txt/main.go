package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

// 采集搜狗哈哈段子
func main() {
	ch := make(chan int)

	startPage := 1
	endPage := 30

	for i := startPage; i <= endPage; i++ {
		go spider(i, ch)
	}

	for i := startPage; i <= endPage; i++ {
		currentPage := <-ch
		fmt.Println("Current page is " + strconv.Itoa(currentPage))
	}
}

// 爬取页面
func spider(page int, ch chan int) {
	url := "http://m.haha.sogou.com/getMore/index?key=text&page=" + strconv.Itoa(page)

	// 获取响应内容
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		fmt.Println("Get url fail.", errs)
		os.Exit(0)
	}

	// 解析内容
	//total := gjson.Parse(body).Get("total").Int()
	//pagenum := gjson.Parse(body).Get("pagenum").Int()
	texts := gjson.Parse(body).Get("list.#.text").Array()
	for _, text := range texts {
		//fmt.Println(text.String() + "\n")

		fl, errs := os.OpenFile("funny.txt", os.O_APPEND|os.O_CREATE, 0644)
		if errs != nil {
			fmt.Println("Open file fail.", err.Error())
			os.Exit(0)
		}

		defer fl.Close()

		bytes := []byte(text.String() + "\n\n")
		fl.Write(bytes)
	}

	ch <- page
}
