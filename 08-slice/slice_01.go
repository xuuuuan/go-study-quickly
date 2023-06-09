// 切片slice（动态数组）
package main

import "fmt"

func newPrintArray(array []int) {
	for index, value := range array {
		fmt.Println("index = ", index, " value = ", value)
	}
	// 会影响main中的数组 这里的动态数组可以理解为引用传递
	array[0] = 100
}

func main() {
	// 动态数组，切片 slice
	myArray := []int{2, 4, 6}

	fmt.Printf("myArray type is %T\n", myArray) // []int

	newPrintArray(myArray)

	fmt.Println(" ======= ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
