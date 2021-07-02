package main

import "fmt"

// 接口示例
type Usb interface {
	Start()
	Stop()
}

type Phone struct{}
type Camera struct{}

func (p *Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p *Phone) Stop() {
	fmt.Println("手机停止工作...")
}
func (c *Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c *Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct{}

func (c *Computer) Worker(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	// computer.Worker(&Phone{})
	computer.Worker(&Camera{})
}
