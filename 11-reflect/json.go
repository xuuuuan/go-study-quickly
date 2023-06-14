package main

import (
	"encoding/json"
	"fmt"
)

type Day struct {
	Date    int    `json:"date"`
	DateStr string `json:"dateStr"`
	Year    int    `json:"year"`
	YearDay int    `json:"yearDay"`
	Week    int    `json:"week"`
	Weekend int    `json:"weekend"`
	Workday int    `json:"workday"`
	Holiday string `json:"holiday"`
}

func main() {
	day := Day{
		Date:    20230614,
		DateStr: "2023-06-14",
		Year:    2023,
		YearDay: 165,
		Week:    3,
		Weekend: 0,
		Workday: 1,
		Holiday: "非节假日",
	}
	// 编码的过程  结构体 ---> json
	// 返回的数据类型是([]byte, error)
	bytes, err := json.Marshal(day)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	jsonStr := string(bytes)
	fmt.Println(jsonStr)

	// 解码的过程 jsonStr ---> []byte ---> 结构体
	var dayObj Day
	err = json.Unmarshal([]byte(jsonStr), &dayObj)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Println(dayObj)
}
