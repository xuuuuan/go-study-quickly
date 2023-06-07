package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

/*
defer会在完成函数体后执行 所以在return后执行
运行结果为
return func called...
defer func called...
*/
func main() {
	deferAndReturn()
}
