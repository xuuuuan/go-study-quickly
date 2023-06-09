// slice的声明
package main

import "fmt"

func main() {
	// 1.声明slice1是一个切片，并且初始化，默认值是1，2，3。 长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 2.声明slice2是一个切片，但是没有分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	// 判断一个切片是否为空
	if slice2 == nil {
		fmt.Println("slice2 是一个空切片")
	} else {
		fmt.Println("slice2 是有空间的")
	}
	// 给slice2开辟3个空间，默认值是0
	slice2 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	// 3.声明slice3是一个切片，同时给slice分配空间，4个空间，初始化值是0
	var slice3 = make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	// 4.声明slice4是一个切片，同时给slice分配空间，4个空间，初始化值是0, 通过:=推导出slice是一个切片
	slice4 := make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

}
