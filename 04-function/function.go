package main

import "fmt"

// 单个返回类型写法
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 返回多个返回值，匿名的
func foo1(a int) (int, int) {
	fmt.Println("----foo1----")
	fmt.Println("a = ", a)
	return 100, 200
}

// 返回多个返回值，有形参名称的
func foo2(a int) (r1 int, r2 int) {
	fmt.Println("----foo2----")
	fmt.Println("a = ", a)

	// r1 r2 属于foo2的形参，  初始化默认的值是0
	// r1 r2 作用域空间 是foo2 整个函数体的{}空间
	fmt.Println("r1 = ", r1, "r2 = ", r2)
	r1 = 300
	r2 = 400
	return
}

func foo3(a int) (r1, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}
func main() {
	fmt.Println(max(3, 8))

	r1, r2 := foo1(1)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	r1, r2 = foo2(2)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	r1, r2 = foo3(3)
	fmt.Println("r1 = ", r1, "r2 = ", r2)
}
