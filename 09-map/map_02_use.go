package main

import "fmt"

func printCityMap(cityMap map[string]string) {
	// 这里的cityMap是引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	fmt.Println("----------------")
}

func main() {
	cityMap := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"UK":    "London",
	}
	printCityMap(cityMap)

	// 增加
	cityMap["US"] = "DC"

	// 删除
	delete(cityMap, "UK")

	// 更改
	cityMap["China"] = "北京"
	printCityMap(cityMap)

	// 查找
	value := cityMap["Japan"]
	fmt.Println("value = ", value)
}
