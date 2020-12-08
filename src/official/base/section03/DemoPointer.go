package main

import "fmt"

func main() {

}

// Go 拥有指针。指针保存了值的内存地址。
//	- 类型 *T: 是指向 T 类型值的指针。其零值为 nil。
//	- & 操作符: 会生成一个指向其操作数的指针。
//	- * 操作符: 表示指针指向的底层值。
// 这也就是通常所说的“间接引用”或“重定向”。
// 与 C 不同，Go 没有指针运算。
func pointerDemo() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
