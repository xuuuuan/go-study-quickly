package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 匿名
	go func(str string, num int) {
		defer fmt.Println("go func defer")
		func() {
			defer fmt.Println("func defer")
			runtime.Goexit()
			fmt.Println("func")
		}()
		fmt.Println("str = ", str, " num = ", num)
	}("mars", 0)

	i := 0
	for {
		i++
	}
}
