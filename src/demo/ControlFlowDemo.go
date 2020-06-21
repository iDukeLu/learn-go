package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	forDemo()
}

var intArray = []int{1, 2, 3, 4, 5, 6}
var a, b, c = 1.0, 2.0, 3.0

// 基本的 for 循环由三部分组成，它们用分号隔开：
//	- 初始化语句：在第一次迭代前执行
//	- 条件表达式：在每次迭代前求值
//	- 后置语句：在每次迭代的结尾执行
// 补充：初始化语句、后置语句 可按需省略
func forDemo() {
	// 0. foreach 循环
	for i, e := range intArray {
		fmt.Printf("下标：%d, 元素：%v \r\n", i, e)
	}

	// 1. 完整 for 循环
	for i := 0; i < 10; i++ {
		fmt.Printf("下标：%d, 元素：%v \r\n", i, intArray[i])
	}

	// 2.省略初始化语句、后置语句的 for 循环
	sum := 0
	for sum < 999 {
		sum++
		fmt.Println(sum)
	}

	// 3. 以上情况，可去掉 ; ，及充当 while 循环
	for sum < 999 {
		sum++
		fmt.Println(sum)
	}

	// 4. 无限 for 循环
	for {
		fmt.Println("无限循环中。。。")
	}
}

// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
// if 可由两部分组成：
//	- 初始化语句：判断前执行，变量作用域 if 范围里
//	- 条件表达式：在每次判断前求值
func ifDemo() {
	// 1. 基本的 if 判断
	if a < b {
		fmt.Print(b)
	}

	// 2. 带有简单初始化语句的 if 判断
	if v := math.Max(a, b); v < c {
		fmt.Print(v)
	}
}

// 除非以 fallthrough 语句结束，否则分支会自动终止。
func switchDemo() {
	// 1. 基本的 switch 选择
	os := runtime.GOOS
	switch os {
	case getDarwin():
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// 2. 带有简单初始化语句的 switch 选择
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// 3. 没有条件的 switch
	// 没有条件的 switch 同 switch true 一样。
	// 这种形式能将一长串 if-then-else 写得更加清晰。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func getDarwin() string {
	return "darwin"
}
