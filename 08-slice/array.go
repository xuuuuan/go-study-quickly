package main

import "fmt"

func printArray(array [4]int) {
	// 这里的array是值拷贝
	// _ 表示匿名的变量 可以不使用
	for _, value := range array {
		fmt.Println("value = ", value)
	}

	// 不会影响main里的array
	array[0] = 100
}

func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray1[0] = 1

	// 初始化固定长度的数组的时候，会按照数据的默认值补全
	myArray2 := [4]int{1, 2, 4}
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, " value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("type of array1 is %T\n", myArray1) // [10]int
	fmt.Printf("type of array2 is %T\n", myArray2) // [4]int

	/*
		因为myArray1的类型是[10]int 不是[4]int 所以会报错
		printArray(myArray1)
	*/
	printArray(myArray2)
	fmt.Println(" ============ ")
	for index, value := range myArray2 {
		fmt.Println("index = ", index, " value = ", value)
	}
}
