package main

import "fmt"

func main() {
	ifDemo(1)
}

// 同 for 一样， if 语句可以在条件表达式前执行一个简单的初始化语句。
func ifDemo(x int) {
	// 条件表达式
	if x < 2 {
		fmt.Print(x)
	}

	// 初始化语句; 条件表达式
	if i := 1; i < x {
		fmt.Print(i)
	}
}
