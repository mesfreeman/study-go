package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
1. 初始化管道
2. 开启协程，爬取资源到管道中
3. 处理管道中的资源数据
4. 监听管道，任务完成则关闭
*/

var (
	chanImageUrls chan string    // 图片链接管道
	chanTask      chan string    // 任务统计
	waitGroup     sync.WaitGroup // 等待组
)

func main() {
	// 1. 初始化管道
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 5)

	// 2. 开启协程，爬取资源到管道中
	totalPage := 1
	for i := 1; i <= totalPage; i++ {
		waitGroup.Add(1)
		// https://www.souutu.com/mbizhi/index_3.html
		go getImgUrls("https://www.souutu.com/mbizhi/index_" + strconv.Itoa(i) + ".html")
	}

	// 任务统计协程
	waitGroup.Add(1)
	go taskSummary(totalPage)

	// 3. 处理管道中的资源数据，即下载图片资源
	for i := 1; i < 5; i++ {
		waitGroup.Add(1)
		go downloadImage()

	}

	// 4. 监听管道，任务完成则关闭
	waitGroup.Wait()
	fmt.Println("爬取任务结束")
}

// 下载图片到本地
func downloadImage() {
	for imageUrl := range chanImageUrls {
		// 处理图片名
		index := strings.LastIndex(imageUrl, "/")
		filename := imageUrl[index+1:]
		filename = strconv.Itoa(int(time.Now().UnixNano())) + "_" + filename

		_, body, errors := gorequest.New().Get(imageUrl).End()
		if errors != nil {
			fmt.Printf("请求图片链接获取内容失败，图片地址：%s \n", imageUrl)
		}

		file := ioutil.WriteFile(filename, []byte(body), 0644)
		if file != nil {
			fmt.Printf("下载图片到本地文件夹中失败，图片地址：%s \n", imageUrl)
		} else {
			fmt.Printf("下载图片到本地文件夹中成功，图片名称：%s \n", filename)
		}
	}
	waitGroup.Done()
}

// 任务统计
func taskSummary(totalPage int) {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("页面【%s】完成了图片采集任务 \n", url)
		count++

		// 任务全部完成，关闭协程
		if count == totalPage {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// 获取指定页面的所有图片链接
func getImgUrls(pageUrl string) {
	_, body, errors := gorequest.New().Get(pageUrl).End()
	if errors != nil {
		fmt.Printf("获取指定页面的所有图片链接失败，页面地址：%s \n", pageUrl)
		panic(0)
	}

	// 正则解析数据
	// lazysrc2x="https://img.souutu.com/2020/1109/20201109015111378.jpg.420.760.jpg 2x"
	subMatches := regexp.MustCompile(`lazysrc2x="(https?:\/\/(.*?).jpg)\s2x`).FindAllStringSubmatch(body, -1)
	fmt.Printf("页面【%s】共找到 %d 个图片资源 \n", pageUrl, len(subMatches))

	// 2. 开启协程，爬取资源到管道中
	for _, item := range subMatches {
		imageUrl := item[1]
		chanImageUrls <- imageUrl
	}

	// 任务统计
	chanTask <- pageUrl
	waitGroup.Done()
}
