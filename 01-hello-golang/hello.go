package main // 包名

import (
	"fmt"
	"time"
)

// main函数
func main() { // '{' 一定得和函数名一行 不然会编译报错
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, World!")
}
