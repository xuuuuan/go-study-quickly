/*
多态
*/
package main

import "fmt"

// Animal 接口不包含任何变量或数据
type Animal interface {
	Sleep()
	GetName() string
}

// Cat 实现接口：要实现一个接口，需要在结构体上实现接口中的所有方法。
// 结构体中的方法必须具有与接口中对应方法的相同名称和签名
type Cat struct {
	Name string
}

// Sleep 地址传递
func (cat *Cat) Sleep() {
	fmt.Println("Cat sleep")
}
func (cat *Cat) GetName() string {
	return cat.Name
}

// --------

type Dog struct {
	Name string
}

// Sleep 值传递
func (dog Dog) Sleep() {
	fmt.Println("Dog sleep")
}
func (dog Dog) GetName() string {
	return dog.Name
}

func ShowAnimal(animal Animal) {
	fmt.Println("----------")
	animal.Sleep()
	fmt.Println(animal.GetName())
}

func main() {
	var animal Animal
	animal = &Cat{"小喵"}
	animal.Sleep()
	fmt.Println(animal.GetName())
	fmt.Println("----------")
	animal = Dog{Name: "阿黄"}
	animal.Sleep()
	fmt.Println(animal.GetName())

	cat := Cat{"little green"}
	dog := Dog{"little yellow"}

	ShowAnimal(&cat)
	ShowAnimal(&dog)
}
