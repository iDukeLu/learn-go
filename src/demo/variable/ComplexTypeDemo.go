package main

import "fmt"

func main() {
	pointDemo()
}

// 指针示例
func pointDemo() {
	i := 1
	// 获取指针
	j := &i
	fmt.Print(j)
}
