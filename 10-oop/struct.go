package main

import "fmt"

// Book 定义一个结构体
type Book struct {
	title  string
	author string
	price  float64
}

func main() {
	var book1 Book
	book1.title = "《Java》"
	book1.author = "Jack"
	book1.price = 29.9
	fmt.Println(book1)

	book2 := Book{
		title:  "《Golang》",
		author: "Mary",
		price:  39.9,
	}
	fmt.Println(book2)

	changeBook(book2)
	fmt.Println(book2)

	changeBookByPointer(&book2)
	fmt.Println(book2)

	book2 = changeBook2(book2)
	fmt.Println(book2)

}

func changeBook(book Book) {
	// 这里是值拷贝传递 所以不会影响main中的book
	book.title = "《JavaScript》"
}

func changeBook2(book Book) Book {
	book.title = "《TypeScript》"
	return book
}

func changeBookByPointer(book *Book) {
	book.title = "《JavaScript》"
}
