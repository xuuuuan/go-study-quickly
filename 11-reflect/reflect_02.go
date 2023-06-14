package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Introduce() {
	fmt.Printf("my name is %s,my age is %d", user.Name, user.Age)
}

func DoFieldsAndMethods(arg interface{}) {
	// 1.获取arg的type和value
	argType := reflect.TypeOf(arg)
	fmt.Println("arg type : ", argType.Name())
	argValue := reflect.ValueOf(arg)
	fmt.Println("arg value : ", argValue)

	// 2.得到字段
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 3.得到方法
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)

	}
}

func main() {
	user := User{1, "Seven", 17}
	DoFieldsAndMethods(user)
}
