package main

import (
	"fmt"
	"reflect"
)

// 全局变量只能用方法1、2、3 方法4是不行的
var globalStr = "go!"

func main() {
	// 1.四种声明方法
	// 1.1 声明一个变量 不同的数据类型有不同的默认值
	var temp string
	fmt.Println("temp = ", temp, "type is ", reflect.TypeOf(temp))

	// 1.2 声明一个变量，初始化一个值
	var name string = "jack"
	fmt.Println("name is ", name)

	// 1.3 在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var a = 10
	fmt.Println("a = ", a, ",type is", reflect.TypeOf(a))

	// 1.4 (常用的方法) 省去var关键字，直接自动匹配
	age := 24
	fmt.Println("age = ", age, ",type is", reflect.TypeOf(age))

	// 2.声明多个变量
	var b, c = 100, "hello"
	fmt.Println("b = ", b, ",c = ", c, "type of b: ", reflect.TypeOf(b), "type of c:", reflect.TypeOf(c))
	var (
		d = 100
		e = true
	)
	fmt.Println("d = ", d, ",e = ", e, "type of d: ", reflect.TypeOf(d), "type of e:", reflect.TypeOf(e))

	fmt.Println(globalStr)
}
