package main

import "fmt"

func main() {
	deferDemo()
}

// defer 语句会将函数推迟到外层函数(一级)返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
// 推迟的函数压入栈时，传人的参数已经被记录
func deferDemo() {
	x := 0
	interFunc(&x)
	fmt.Println("deferDemo")
}

func interFunc(x *int) {
	defer deferFunc(x)
	fmt.Printf("interFunc: %v \n", *x+1)
}

func deferFunc(x *int) {
	fmt.Printf("deferFunc: %v \n", *x+1)
}
