package main

import "fmt"

func main() {
	mapDemo4()
}

// 映射的零值为 nil。nil 映射既没有键，也不能添加键。
// make 函数会返回给定类型的映射，并将其初始化备用。
func mapDemo1() {
	var m map[string]string
	fmt.Printf("nil 映射: %v \n", m)

	m = make(map[string]string)
	m["key"] = "value"
	fmt.Printf("映射: %v \n", m)
}

// 映射的文法与结构体相似，不过必须有键名。
func mapDemo2() {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	fmt.Printf("映射: %v \n", m)

}

// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
func mapDemo3() {
	var m = map[string]Vertex1{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Printf("映射: %v \n", m)
}

type Vertex1 struct {
	X, Y int
}

// map 常用操作
func mapDemo4() {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// 获取元素
	fmt.Printf("获取 key1 的值: %v \n", m["key1"])

	// 删除元素
	delete(m, "key2")
	fmt.Printf("删除 key2 的映射: %v \n", m)

	// 检测某个键是否存在
	_, ok := m["key1"]
	fmt.Printf("是否存在 key1: %v \n", ok)
}
