package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型、基本数据类型转换
func main() {
}

// 基本数据：整数型、浮点型、复数型、字符型、字符串型
// 	- 整数型
// 		- 有符号：int int8 int16 int32 int64 rune
// 		- 无符号：uint uint8 uint16 uint32 uint64 uintptr byte
// 	- 浮点型：float32 float64
// 	- 复数型：complex64 complex128
// 	- 字符型：byte
// 	- 字符串型：string
func baseTypeDemo() {
	i := 8
	var i8 int8 = 127
	fmt.Printf("%T", i)
	fmt.Printf("%T", i8)
}

func numConvDemo() {

}

func strConvDemo() {
	intStr := fmt.Sprintf("%d", 99)
	//floatStr := fmt.Sprintf("%f", 99.9)
	floatStr1 := fmt.Sprintf("%v", 99.9)
	boolStr := fmt.Sprintf("%v", true)
	fmt.Printf("intStr: %v, floatStr: %v, boolStr: %v", intStr, floatStr1, boolStr)

	// 整数、进制
	//i := strconv.FormatInt(99, 10)
	//i1 := strconv.Itoa(99)
	//
	// 浮点数、格式标记（b：二进制指数、e：十进制指数、E：十进制指数、f：没有指数、g：大指数、G：大指数）、精度、浮点类型
	f := strconv.FormatFloat(99.9, 'f', 2, 64)
	//
	//b := strconv.FormatBool(true)

	fmt.Printf("i: %s", f)

}
