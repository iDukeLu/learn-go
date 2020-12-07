package main

import "fmt"

func main() {
	forDemo(7)
}

// Go 只有一种循环结构：for 循环。
// 基本的 for 循环由三部分组成，它们用分号隔开：（初始化语句和后置语句是可选的）
//	- 初始化语句：在第一次迭代前执行
//	- 条件表达式：在每次迭代前求值
//	- 后置语句：在每次迭代的结尾执行
func forDemo(x int) {
	//
	//初始化语句; 条件表达式; 后置语句
	for i := 0; i < x; i++ {

	}

	// 条件表达式
	for x < 100 {

	}

	// 无限循环
	for {
		fmt.Print(x)
	}
}
