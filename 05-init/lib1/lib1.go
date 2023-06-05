package lib1

import "fmt"

// Max 首字母M大写代表是对外暴露的function
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 初始化函数被其他文件import导入的时候就会执行
func init() {
	fmt.Println("lib1 init...")
}
