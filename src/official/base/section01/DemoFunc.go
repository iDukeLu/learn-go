package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

var (
	a, b int
)

// Go 方法有以下几个特点：
// 	- 相同类型的 参数列表 和 返回列表 可省略类型
// 	- 可拥有多个返回值
// 	- 可对返回值命名
func main() {
	fmt.Println(funcDemo4())
}

func funcDemo1(x int, y int) int {
	return x + y
}

// 相同类型的参数可省略类型
func funcDemo2(x, y int) int {
	return x + y
}

// 可拥有多个返回值
func funcDemo3(x, y int) (int, int) {
	return x, y
}

// 可对返回值命名
func funcDemo4() (x, y int) {
	return x, y // 或 return（会自动返回命名相同的变量）
}
