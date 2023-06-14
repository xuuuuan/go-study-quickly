package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	// 标签的格式为`key:"value"`
	Name  string `info:"name" doc:"名称"`
	Price int    `info:"price" doc:"价格"`
}

func findTag(arg interface{}) {
	// 因为不确定arg里的字段是否有值 所以需要使用 Elem()
	elem := reflect.TypeOf(arg).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fmt.Printf("tag-info: %s, tag-doc: %s\n", field.Tag.Get("info"), field.Tag.Get("doc"))
	}
}

func main() {
	var order Order
	findTag(&order)
}
