/*
类的继承
*/
package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (human *Human) Walk() {
	fmt.Println("human walk")
}

func (human *Human) Eat() {
	fmt.Println("human eat")
}

// Superman 类继承Human类
type Superman struct {
	Human
	level int
}

// Walk 方法的重写
func (superman *Superman) Walk() {
	fmt.Println("superman walk")
}

// Fly 子类新方法
func (superman *Superman) Fly() {
	fmt.Println("superman fly")
}

func main() {
	human := Human{"jack", 17}
	human.Eat()
	human.Walk()
	fmt.Println("-------------")
	//superman := Superman{
	//	Human{"jacky", 18}, 88,
	//}
	var superman Superman
	superman.name = "jacky"
	superman.age = 18
	superman.level = 10

	superman.Eat()
	superman.Walk()
	superman.Fly()
}
