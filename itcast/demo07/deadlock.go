package main

// 死锁的三种形式

// 第一种：单 go 程自己死锁
// channel 应该在至少2个以上的 go 程中进行通信，否则死锁！！！
/*
func main() {
	ch := make(chan int)

	ch <- 67
	num := <-ch
	fmt.Println(num)
}
*/

// 第二种：go 程间 channel 访问顺序导致的死锁
// 使用 channel 一端读，要保证另一端写
/*
func main() {
	ch := make(chan int)

	num := <-ch
	go func() {
		ch <- 76
	}()
	fmt.Println(num)
}
*/

// 第三种：多 channel 交叉导致的死锁
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}

}
