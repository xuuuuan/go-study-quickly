package main

import (
	"fmt"
	"reflect"
)

func MyPrint(arg interface{}) {
	fmt.Println("reflect type = ", reflect.TypeOf(arg))
	fmt.Println("reflect value = ", reflect.ValueOf(arg))
}

func main() {
	var num = 3.1415926
	MyPrint(num)
}
