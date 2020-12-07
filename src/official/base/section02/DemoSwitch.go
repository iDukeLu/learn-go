package main

import "fmt"

func main() {
	switchDemo(7)
}

// switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。
// Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。case 可以是：常量、变量、方法。
func switchDemo(x int) {

	// 初始语句; 条件
	// case 可以是：常量、变量、方法
	switch y := 2; x {
	case case1(1):
		fmt.Print(1)
	case y:
		fmt.Print(2)
	case 3:
		fmt.Print(3)
	default:
		fmt.Print(4)
	}

	// 条件
	switch x {
	case case1(1):
		fmt.Print(1)
	case x:
		fmt.Print(2)
	case 3:
		fmt.Print(3)
	default:
		fmt.Print(4)
	}

	// 无条件的 switch 同 switch true 一样。
	switch {
	case x < 1:
		fmt.Print(1)
	case x > 1:
		fmt.Print(2)
	default:
		fmt.Print(4)
	}
}

func case1(x int) int {
	return x
}
