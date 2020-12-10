package main

import "fmt"

var (
	s string
	i = 1
)

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
// 	- 指定初始值：变量声明可以包含初始值，每个变量对应一个。
// 	- 省略类型：如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
// 	- 短变量：在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
func main() {
	var f bool

	var f1 float64 = 1
	var f2 = 1

	f3 := 1

	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(f)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
}
