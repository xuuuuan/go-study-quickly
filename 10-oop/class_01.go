/*
类的定义
*/
package main

import "fmt"

// Hero 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}

func (hero *Hero) PrintHero() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("level = ", hero.level)
}

func (hero *Hero) SetName(newName string) {
	// 想要改变对象的属性 需要使用地址引用
	hero.Name = newName
}

func main() {
	hero := Hero{
		Name:  "li bai",
		Ad:    680,
		level: 15,
	}
	hero.PrintHero()
	hero.SetName("han xin")
	hero.PrintHero()
}
