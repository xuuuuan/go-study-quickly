package main // 包名

import (
	"fmt"
	"time"
)

// main函数
func main() { // '{' 一定得和函数名一行 不然会编译报错
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, World!")

	currentTime := time.Now()
	// 自定义输出格式
	// time.DateTime = "2006-01-02 15:04:05"
	formattedTime := currentTime.Format(time.DateTime)

	fmt.Println("当前时间:", formattedTime)
}
