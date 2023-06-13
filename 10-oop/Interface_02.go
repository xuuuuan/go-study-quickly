/*
空接口
*/
package main

import "fmt"

func myPrint(arg interface{}) {
	// 断言是string类型
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg type is string, value = ", value)
	} else {
		fmt.Println("arg type is not string")
	}

	switch arg.(type) {
	case float64:
		fmt.Println("arg type is float64")
	}
}

func main() {
	myPrint("abc")
	myPrint(123)
	myPrint(3.14)
}
