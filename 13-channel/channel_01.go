/*
无缓冲channel
*/
package main

import "fmt"

func main() {
	// 定义一个channel
	channel := make(chan int)
	go func() {
		defer fmt.Println("sub goroutine exit")
		fmt.Println("sub goroutine is running")
		// 将666发送给channel
		channel <- 666
	}()
	// 从channel中接受数据赋值给num
	num := <-channel
	fmt.Println("num = ", num)
	defer fmt.Println("main goroutine exit")
}
