package main

import "fmt"

func main() {
	// 要注意的是，map是无序的，因此不能像数组或切片一样通过索引访问元素
	// ===> 第一种声明方式
	// map[key] value
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("map is null")
	}
	printMap(myMap1)
	myMap1 = make(map[string]string, 3)

	myMap1["one"] = "java"
	myMap1["two"] = "MySQL"
	myMap1["three"] = "js"

	for key, value := range myMap1 {
		fmt.Println("key = ", key, "value = ", value)
	}
	printMap(myMap1)

	// ===> 第二种声明方式
	myMap2 := make(map[string]string)
	printMap(myMap2)
	myMap2["1"] = "java"
	myMap2["2"] = "c++"
	myMap2["3"] = "python"
	printMap(myMap2)

	//===> 第三种声明方式
	myMap3 := map[string]string{
		"one":   "vue",
		"two":   "c++",
		"three": "python",
	}
	printMap(myMap3)

}

func printMap(myMap map[string]string) {
	fmt.Printf("map = %v, len = %d\n", myMap, len(myMap))
}
