/*
有缓冲channel
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 3)
	go func() {
		defer fmt.Println("sub goroutine exit")
		for i := 0; i < 4; i++ {
			channel <- i
			fmt.Println("i = ", i, " len = ", len(channel), " cap = ", cap(channel))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-channel
		fmt.Println("num = ", num)
	}
	fmt.Println("main goroutine exit")
}
