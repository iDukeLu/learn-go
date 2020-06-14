package main

import "fmt"

func main() {

}

// 变量的三种声明方式
func varDemo() {
	var i int
	var j = 1
	k := 3
	fmt.Printf("无赋值声明：%v", i)
	fmt.Printf("自动类型检测：%v", j)
	fmt.Printf("短声明，仅能在函数中使用：%v", k)
}
