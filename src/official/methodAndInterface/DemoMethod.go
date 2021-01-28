package main

import (
	"fmt"
)

func main() {
	methodDemo4()
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func methodDemo1() {
	user := new(User)
	*user = user.setNameAndReturn("张三")
	fmt.Println(user)
}

type User struct {
	name string
	age  int
}

func (user User) getName() string {
	return user.name
}

func (user User) setNameAndReturn(name string) User {
	user.name = name
	return user
}

// 你也可以为非结构体类型声明方法。
func methodDemo2() {
	var i Integer = -10
	fmt.Println(i.abs())
}

type Integer int

func (i Integer) abs() Integer {
	if i < 0 {
		return -i
	}
	return 1
}

// 你可以为指针接收者声明方法。
// 这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）
// 指针接收者的方法可以修改接收者指向的值，若使用值接收者，那么方法会对原始值的副本进行操作
// 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用
func methodDemo3() {
	user := new(User)
	user.setName("张三")
	fmt.Println(user)
}

func (user *User) setName(name string) {
	user.name = name
}

// 方法与指针重定向
// 以指针为参数的函数，必须接受一个指针
// 以指针为接收者的方法，被调用时，接收者既能为值又能为指针
// 以值为参数的函数，必须接受一个值
// 以值为接收者的方法，被调用时，接收者既能为值又能为指针
// 使用指针接收者的原因有二：
//	首先，方法能够修改其接收者指向的值。
//	其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
func methodDemo4() {

}

func func1() {

}
