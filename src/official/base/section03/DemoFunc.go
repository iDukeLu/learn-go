package main

import (
	"fmt"
	"math"
)

func main() {

}

// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。
func funcDemo0() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
// 例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
func funcDemo1() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
