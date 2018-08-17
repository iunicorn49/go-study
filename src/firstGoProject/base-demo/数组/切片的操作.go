package main

import "fmt"

// 对切片进行操作会影响原始数组
// 如果对切片的操作, 添加数据后, 超过了原始数组长度, go语言内部会生成一个新的数组, 这个切片会view那个新的数组

func sliceAppend () { // 添加
	arr := [...]int{0,1,2,3,4,5}
	s1 := arr[2:4]
	s2 := append(s1, 10)
	s3 := append(s2, 11,12)
	fmt.Println(s2) // [2 3 10] // 会影响原始数组
	fmt.Println(s3) // [2 3 10 11 12] // 因操作slice而超出原始数组的话, 会重新view一个新的数组
	fmt.Println(arr) // [0 1 2 3 10 5]
}

func sliceops () { // 创建
	var s []int // 创建了slice但是没有view任何数组, 会有一个默认值 nil, 长度为0
	for i := 0; i < 5; i++ {
		s = append(s, 2 * i)
	}
	fmt.Println(s) // [0 2 4 6 8]
	s2 := make([]int, 4, 8) // 创建一个[0 0 0 0] len 为 4, cap 为 8 的slice,
	fmt.Println(s2) // [0 0 0 0]
	copy(s2, s) // 拷贝
	fmt.Println(s2) // [0 2 4 6]
	s2 = append(s2[:1], s2[2:]...) // 删掉下标为 1 的元素, 没有内建方法, 可以自己用 append 拼一个新的slice
	fmt.Println(s2) // [0 4 6]
}

func delete (s[]int, i int) {
	s = append(s[:i], s[i + 1:]...)
	fmt.Println(s) // [0 4 6 8]
}

func main () {
	//sliceAppend()
	//sliceops()
	var s1 []int
	for i := 0; i < 5; i++ {
		s1 = append(s1, i * 2)
	}
	fmt.Println("s1 =", s1) // [0 2 4 6 8]
	delete(s1, 1)	// [0 4 6 8]
	fmt.Println(s1) // [0 4 6 8 8]
}
