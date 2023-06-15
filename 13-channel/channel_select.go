package main

import "fmt"

func count(c1 chan int, c2 chan int) {
	x := 0
	for {
		// 使用select监听多个channel
		select {
		case c1 <- x:
			// 如果c1可写
			x++
		case <-c2:
			// 如果c2可读
			fmt.Println("count quit")
			return
		}
	}
}

func main() {
	channel := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()

	// main go
	count(channel, quit)
}
