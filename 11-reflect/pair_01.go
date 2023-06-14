package main

import (
	"fmt"
)

func main() {
	// pair<type: string, value: >
	var str string
	// pair<type: string, value: go>
	str = "go"

	var all interface{}
	// pair<type: string, value: go> 会一直传递下去
	all = str

	value, _ := all.(string)
	fmt.Println(value)
}
