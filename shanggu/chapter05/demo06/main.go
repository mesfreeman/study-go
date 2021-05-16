package main

import "fmt"

func main() {
	var w byte
	fmt.Println("请输入一个字符：")
	fmt.Scanf("%c", &w)

	switch w {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	default:
		fmt.Println("输入不合法~")
	}

	var n1 int32 = 20
	var n2 int32 = 10
	switch n1 {
	case n2, 20: // case后面可以有多个表达式
		fmt.Println("ok1")
	default:
		fmt.Println("ok2")
	}

	var num = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // switch 穿透
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("...")
	}

	var x interface{}
	var y = 10.1
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\n", i)
	case int:
		fmt.Println("x 是 int 型")
	case float32:
		fmt.Println("x 是 float32 型")
	case float64:
		fmt.Println("x 是 float64 型")
	default:
		fmt.Println("x 是不知型")
	}
}
