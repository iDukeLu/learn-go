package main

import "fmt"

func main() {
	sliceDemo9()
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔。
// 它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
func sliceDemo0() {
	arr := [7]int{1, 2, 3, 4}
	slice := arr[0:4]
	fmt.Printf("切片：%v \n", slice)
}

// 更改切片的元素会修改其底层数组中对应的元素。
// 切片并不存储任何数据，它只是描述了底层数组中的一段。
func sliceDemo1() {
	arr := [7]int{1, 2, 3, 4}
	fmt.Printf("修改前数组：%v \n", arr)

	// 更改切片的元素会修改其底层数组中对应的元素。
	slice := arr[0:4]
	slice[0] = 0
	fmt.Printf("修改切片后的数组：%v \n", arr)
}

// 切片文法类似于没有长度的数组文法。
func sliceDemo2() {
	arr := [7]int{1, 2, 3, 4}
	fmt.Printf("数组：%v", arr)

	// 切片文法类似于没有长度的数组文法。
	// 下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片
	slice := []int{1, 2, 3, 4}
	fmt.Printf("切片：%v", slice)
}

// 在进行切片时，你可以利用它的默认行为来忽略上下界。
// 下界默认值：0，上界默认值：切片的长度。
func sliceDemo4() {
	arr := [7]int{1, 2, 3, 4}
	fmt.Printf("数组：%v", arr)

	// 切片下界的默认值为 0，上界则是该切片的长度。
	slice := arr[:]
	fmt.Printf("切片：%v", slice)
}

// 切片拥有 长度 和 容量。
// 切片的长度：包含元素的个数
// 切片的容量：数组的长度
// 长度获取：len(s)、容量获取：cap(s)
// 你可以通过重新切片来扩展一个切片，给它提供足够的容量
func sliceDemo5() {

	slice := []int{1, 2, 3, 4}
	slice = slice[:0]
	fmt.Printf("原切片长度：%v，原切片容量：%v \n", len(slice), cap(slice))

	// 拓展其长度
	slice = slice[:4]
	fmt.Printf("扩展后切片长度：%v，扩展后切片容量：%v \n", len(slice), cap(slice))

	// 减少容量
	slice = slice[2:]
	fmt.Printf("减容后切片长度：%v，减容后切片容量：%v \n", len(slice), cap(slice))
}

// 切片的零值是 nil。
// nil 切片的长度和容量为 0 且没有底层数组。
func sliceDemo6() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// 切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
func sliceDemo7() {
	slice := make([]int, 0)
	fmt.Println(slice, len(slice), cap(slice))
}

// 切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。
func sliceDemo8() {
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice, len(slice), cap(slice))
}

// 切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。
func sliceDemo9() {
	arr := [1]int{1}
	slice := arr[:]

	// 切片添加时，会自动扩容，扩容基于数组复制实现
	slice = append(slice, 1, 2, 3)

	// 赋值，原数组未改变，原因：切片已经在 append() 时扩展复制返回新的数组
	slice[0] = 7
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(slice, len(slice), cap(slice))
}
