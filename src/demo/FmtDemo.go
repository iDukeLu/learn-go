package main

import "fmt"

// 打印：print、println、printf
// 键盘录入：scanln、scanf
func main() {

}

func printDemo() {
	fmt.Print("hello world! 我是不换行的打印")
	fmt.Println("hello world! 我是要换行的打印")
}

// 格式化打印示例
// %v	按值的本来值输出
// %+v	在 %v 基础上，对结构体字段名和值进行展开
// %#v	输出 Go 语言语法格式的值
// %T	输出 Go 语言语法格式的类型和值
// %%	输出 % 本体
// %d	整型以十进制方式显示
// %f	浮点数
// %s	字符串
// %U	Unicode 字符
// %p	指针，十六进制方式显示
func printfDemo() {

}
