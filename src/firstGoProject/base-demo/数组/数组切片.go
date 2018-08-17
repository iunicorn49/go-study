package main

import "fmt"

// Slice(切片)
// arr[:] 完整视图
// arr[:2] 第0位到第1位
// arr[2:6] 第2位到第5位
// arr[2:] 第2位到最后

func updateSlice(s []int) {
	s[0] = 100
}

func reslice() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:6]
	// 在切片上重新slice的时候, 如果不够, 会从原始数组上继续往后读, 所会出现, 二次slice范围可以超出第一个slcie的情况
	fmt.Println(s2) // [5 6 7]
}

func main() {
	//arr := [...]int{0,1,2,3,4,5,6,7}

	//updateSlice(arr[2:]) // 传入的是数组的视图, 会影响原始数组
	//fmt.Println(arr) // [0 1 100 3 4 5 6 7]

	// 多次 slice
	//s1 := arr[2:6]
	//fmt.Println(s1)
	//s2 := s1[1:3]
	//fmt.Println(s2)
	//updateSlice(s2) // 多次 切片后, 还是是影响原始数组
	//fmt.Println(arr) // [0 1 2 100 4 5 6 7]

	reslice()
}