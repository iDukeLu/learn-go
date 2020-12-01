package main

import "fmt"

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%+v", v)
	fmt.Printf("获取结构体的属性：%v", v.X)

	p := &v
	fmt.Printf("获取结构体的属性：%v", p.X)

}

// 一个结构体（struct）就是一组字段（field）。
type Vertex struct {
	X int
	Y int
}
