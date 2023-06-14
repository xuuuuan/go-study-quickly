package main

import (
	"fmt"
	"time"
)

// 子goroutine
func task() {
	i := 0
	for {
		i++
		fmt.Println("task i = ", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 创建一个go程 去执行task() 流程
	go task()

	i := 0
	for {
		if i == 10 {
			// 当主goroutine终止，子goroutine也会终止
			break
		}
		i++
		fmt.Println("main i = ", i)
		time.Sleep(1 * time.Second)
	}
}
