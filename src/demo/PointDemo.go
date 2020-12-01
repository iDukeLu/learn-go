package main

import "fmt"

func main() {

}

func pointDemo() {
	// 类型 *T 是指向 T 类型值的指针。其零值为 nil。
	var p *int

	// & 操作符会生成一个指向其操作数的指针。
	i := 42
	p = &i

	// * 操作符表示指针指向的底层值。这也就是通常所说的“间接引用”或“重定向”。
	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 21         // 通过指针 p 设置 i
}
