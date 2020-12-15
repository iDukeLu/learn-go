package main

import "fmt"

func main() {
	methodDemo()
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func methodDemo() {
	user := new(User)
	user.setName("张三")
	fmt.Println(user)
}

type User struct {
	name string
	age  int
}

func (user User) getName() string {
	return user.name
}

func (user User) setName(name string) {
	user.name = name
}
