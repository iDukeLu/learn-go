package main

import "fmt"

// defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
// 即，defer 修饰的语句或函数会立即求值，但在外层函数返回之后才执行调用

// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
// 即，多个 defer 执行时，按倒序执行
func main() {

	defer sayWorld1()
	defer sayWorld2()
	fmt.Println("hello ")
}

func sayWorld1() {
	fmt.Println("world1")
}

func sayWorld2() {
	fmt.Println("world2")
}
