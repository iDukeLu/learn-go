package main

import (
	"fmt"
	"strconv"
)

func main() {
	baseTypeConvDemo()
}

// 基本数据：整数型、浮点型、复数型、字符型、字符串型
// 整数型
// 	- 有符号：int int8 int16 int32 int64 rune
// 	- 无符号：uint uint8 uint16 uint32 uint64 uintptr byte
// 浮点型：float32 float64
// 复数型：complex64 complex128
// 字符型：byte
// 字符串型：string
func baseTypeDemo() {
	i := 8
	var i8 int8 = 127
	fmt.Printf("%T", i)
	fmt.Printf("%T", i8)
}

// %v	按值的本来值输出
// %+v	在 %v 基础上，对结构体字段名和值进行展开
// %#v	输出 Go 语言语法格式的值
// %T	输出 Go 语言语法格式的类型和值
// %%	输出 % 本体
// %d	整型以十进制方式显示
// %U	Unicode 字符
// %f	浮点数
// %p	指针，十六进制方式显示
func baseTypeConvDemo() {

	intStr := fmt.Sprintf("%d", 99)
	//floatStr := fmt.Sprintf("%f", 99.9)
	floatStr1 := fmt.Sprintf("%v", 99.9)
	boolStr := fmt.Sprintf("%v", true)
	fmt.Printf("intStr: %v, floatStr: %v, boolStr: %v", intStr, floatStr1, boolStr)

}

func convStr() {
	i0 := strconv.Itoa(99)
	b := strconv.FormatBool(true)

	// 整数、进制
	i := strconv.FormatInt(99, 10)

	// 浮点数、格式标记（b：二进制指数、e：十进制指数、E：十进制指数、f：没有指数、g：大指数、G：大指数）、精度、浮点类型
	f := strconv.FormatFloat(99.9, 'f', 10, 64)

}
