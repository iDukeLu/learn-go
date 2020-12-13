package main

import "fmt"

func main() {
	rangeDemo0()
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// for 循环的 range 形式可遍历切片或映射。
// 当使用 for 循环遍历切片时，每次迭代都会返回两个值。
// 第一个值: 当前元素的下标，第二个值: 该下标所对应元素的一份副本。
func rangeDemo0() {
	for i, v := range pow {
		fmt.Printf("arr[%v] = %v \n", i, v)
	}
}
