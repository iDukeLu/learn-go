package main

import "fmt"

func main() {
	fmt.Println(add1(42, 13))
}

func add1(x int, y int) int {
	return x + y
}

// 相同类型的参数可省略类型
func add2(x, y int) int {
	return x + y
}

// 可拥有多个返回值
func swap(x, y int) (int, int) {
	return x, y
}

// 可对返回值命名
func split1() (x, y int) {
	x = x + 1
	y = y - 1
	return x, y // 或 return（会自动返回命名相同的变量）
}
