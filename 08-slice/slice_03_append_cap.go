package main

import "fmt"

func main() {

	// 初始化一个切片 长度为3 容量为5
	numbers := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)

	// 重点！！！
	// 向一个容量cap已经满的slice 追加元素，cap = cap * 2 = 10  下一次cap满的时候会2*10=20
	// 所以 cap = 5 -> 10 -> 20 -> 40...
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("---------")
	var numbers2 = make([]int, 3)
	// 没有指定cap 就默认 cap = len =3
	// len = 3, cap = 3, slice = [0 0 0]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	// len = 4, cap = 6, slice = [0 0 0 1]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
