package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

// 协程综合练习实例
var (
	totalPage = 135
	hostName  = "https://qq.yh31.com"
)

type ImageInfo struct {
	Title string
	Url   string
}

// 抓取图片页面链接
func fetchUrl(urlChan chan<- string) {
	for i := totalPage; i > 1; i-- {
		url := "https://qq.yh31.com/ka/zr/List_" + strconv.Itoa(i) + ".html"
		urlChan <- url
		fmt.Println("采集页面链接：", url)
	}
	close(urlChan)
}

// 抓取图片链接
func fetchImgUrl(urlChan <-chan string, imgChan chan<- ImageInfo, exitChan chan<- bool) {
	for {
		url, ok := <-urlChan
		if !ok {
			exitChan <- true
			break
		}

		// 获取网页源代码
		_, body, err := gorequest.New().
			Proxy("http://14.116.213.100:8081").
			Timeout(time.Second * 5).
			Get(url).End()
		if err != nil {
			fmt.Println("出错啦", url)
			return
		}

		// 正则图片地址
		subMatches := regexp.MustCompile(`target="_blank"><img src="(.*?\.gif)" alt="(.*?)" border="0"\/>`).FindAllStringSubmatch(body, -1)
		fmt.Printf("页面【%s】共找到 %d 个图片资源 \n", url, len(subMatches))

		for _, item := range subMatches {
			if item[1] == "" || item[2] == "" {
				continue
			}

			imgInfo := ImageInfo{
				Title: item[2],
				Url:   hostName + strings.Replace(item[1], "_1.gif", ".gif", 1),
			}
			imgChan <- imgInfo
		}
	}
}

func main() {
	urlChan := make(chan string, 6)
	imgChan := make(chan ImageInfo, 10000)
	exitChan := make(chan bool, 3)

	go fetchUrl(urlChan)
	for i := 0; i < 3; i++ {
		go fetchImgUrl(urlChan, imgChan, exitChan)
	}

	go func() {
		for i := 0; i < 3; i++ {
			<-exitChan
		}

		close(imgChan)
		close(exitChan)
	}()
	filename := "/Users/double/projects/study-go/shanggu/chapter16/test01/img.txt"
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	for imgInfo := range imgChan {
		fmt.Printf("url:%v\ttitle:%v\n\n", imgInfo.Url, imgInfo.Title)
		sqlStr := fmt.Sprintf("INSERT INTO `funnies` (`title`, `cover`, `intro`, `isTop`, `type`, `tags`, `size`, `url`, `m3u8Url`, `content`, `ext`, `createdAt`, `updatedAt`) VALUES ('%v', '', '', 0, 'gif', '动态图', 87185, '%v', '', '', '{\"wxSyncStatus\": true}', 1625995695, 1625995695);\n", imgInfo.Title, imgInfo.Url)
		file.WriteString(sqlStr)
	}
}
