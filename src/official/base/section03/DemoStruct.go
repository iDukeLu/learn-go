package main

import "fmt"

// 一个结构体（struct）就是一组字段（field）。
type User struct {
	Name   string
	Gender string
	Age    int
	Phone  string
}

type Vertex struct {
	X, Y int
}

// 结构体文法通过直接列出字段的值来新分配一个结构体。
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	u := User{}

	// 结构体字段使用点号来访问。
	fmt.Print(u.Name)

	// 结构体字段允许我们使用隐式间接引用，直接写 p.X 就可以。
	fmt.Print(&u.Name)
}
