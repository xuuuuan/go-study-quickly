package main

import "fmt"

func main() {
	a := 1
	changeValue(a)
	fmt.Println("a =", a)

	changeValueByPointer(&a)
	fmt.Println("a =", a)

	b := 20
	swap(a, b)
	fmt.Println("a =", a, " b =", b)
	swapByPointer(&a, &b)
	fmt.Println("a =", a, " b =", b)
}

// 并不会影响main函数中的p
func changeValue(p int) {
	p = 10
}

func changeValueByPointer(p *int) {
	*p = 10
}

// 并不会影响main函数中的a,b
func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swapByPointer(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
