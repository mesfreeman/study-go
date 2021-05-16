package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入小臭臭的成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一辆BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iphone7 plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一台ipad")
	} else {
		fmt.Println("奖励一顿打")
	}
}
