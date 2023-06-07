package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

/*
1.defer的执行顺序
defer会在函数体执行完了再执行
执行顺序: defer是压栈 先进后出所以下面三个defer函数的执行顺序是3->2->1
所以输出结果为
D
C
B
A
*/
func main() {
	defer func1()
	defer func2()
	defer func3()
	fmt.Println("D")
}
