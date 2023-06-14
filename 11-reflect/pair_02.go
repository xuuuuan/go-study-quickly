package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

// Book 具体类型
type Book struct {
	Name string
}

func (b *Book) Read() {
	fmt.Println("read book")
}

func (b *Book) Write() {
	fmt.Println("write book")
}

func main() {
	// book: pair<type:Book, value:book{"<ABC>"}地址>
	book := &Book{"<ABC>"}

	// reader: pair<type:, value:>
	var reader Reader
	// reader: pair<type:Book, value:book{"<ABC>"}地址>
	reader = book
	reader.Read()

	var writer Writer
	// writer: pair<type:Book, value:book{"<ABC>"}地址>
	writer = reader.(Writer) //此处的断言为什么会成功？ 因为writer reader 具体的type是一致. 可以参考java中的强转
	writer.Write()
}
