package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func handleErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(0)
	}
}

func main() {
	var (
		rePhone = `1[23456789]\d{9}`
	)

	url := "https://www.haomagujia.com/"
	resp, err := http.Get(url)
	handleErr(err, "go get "+url)

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	// 正则手机号
	phones := regexp.MustCompile(rePhone).FindAllStringSubmatch(html, -1)
	for _, item := range phones {
		fmt.Println(item[0])
	}
}
