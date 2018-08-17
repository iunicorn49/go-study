package main

import "fmt"

// 长度不同的数组被认为是不同的数据类型
func basicArr() {
	var arr1 [5]int // 定义一个有5个int的数组,不给初始值的情况下,里面用 0 来填充
	arr2 := [3]int{1,3,5} // := 的定义方式需要初始值
	arr3 := [...]int{2,4,6,8,10} // 和上面相似, 只不过这次数组数量 用 '...' 可以让编译器自己数

	fmt.Println(arr1,arr2,arr3)
}

func twoDimension() { // 二维数组
	var grid [4][5]bool // 4行5列的数组
	fmt.Println(grid)
}

func iteration() { // 遍历数组
	list := [...]string{"a", "b", "c", "d"}

	// 数组用 range 进行迭代, i 为下标, v 为值, 如果只想要 v ,可以将 i 替换成 '_'
	for i, v := range list {
		fmt.Println(i, v)
	}

	for _, v := range list {
		fmt.Println(v)
	}
}

// 数组是值类型, 传参会把数组拷贝一份, 不会影响原始数组
// arr[] 代表一个数组的切片视图
// arr[5] 才是真正的数组
func pringtArray(arr [5]string) {
	arr[0] = "ZZZ" // 仅仅在这个作用域内 arr[0] 的值发生改变, 外部再打印, 还是原来的数组
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println(arr)
}

// 用指针的情况下, 会影响到原始数组
func pringtArrayFromPointer(arr *[5]string) {
	arr[0] = "这"
	arr[1] = "来"
	arr[2] = "自"
	arr[3] = "指"
	arr[4] = "针"
}

func main() {
	//basicArr()
	//twoDimension()
	//iteration()
	arr := [5]string{"a", "b", "c", "d", "e"}
	//pringtArray(arr)
	//fmt.Println(arr) // 上面函数执行完以后, 不会影响外面的 arr
	pringtArrayFromPointer(&arr) // 用指针传入, 会影响原始数组
	fmt.Println(arr)
}
