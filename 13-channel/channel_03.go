package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		// close 可以关闭一个channel
		close(channel)
	}()

	/*	for {
		// 先执行 ; 前的语句 判断条件为 ; 后
		if data, ok := <-channel; ok {
			fmt.Println("data = ", data)
		} else {
			fmt.Println("no data")
			break
		}
	}*/
	// 上面的代码可以简写为
	for data := range channel {
		fmt.Println("data = ", data)
	}

	fmt.Println("main goroutine exit")
}
