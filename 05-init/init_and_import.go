package main

import (
	"fmt"
	// 匿名导包，仅仅执行包内的init函数，不需要执行里面的其他函数
	_ "go-study-quickly/05-init/lib1"
	// 起别名
	myLib2 "go-study-quickly/05-init/lib2"
)

func main() {
	//fmt.Println(lib1.Max(10, 6))
	fmt.Println(myLib2.Min(10, 6))
}
