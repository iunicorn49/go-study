package main

import "fmt"

// 注意:
// 1. map中最后项也要用 "," 结尾
// map[key的类型]val的类型
// map使用哈希表, 必须可以比较相等
// 除了 slice, map, function 的内建类型都可以作为 key
// Struct类型不包含上述字段, 也可以作为 key

func createMap () { // 创建
	// 直接声明
	m := map[string]string {
		"name": "ATOM",
		"age": "23",
	}
	fmt.Println(m)

	// make 的方式
	mapFromMake := make(map[string]int) // empty map
	var mapFromVar map[string]int // nil
	// nil 和 empty map 都是可以参与运算的, 没什么区别
	fmt.Println(mapFromMake, mapFromVar)
}

func iteration () { // map 的迭代
	m := map[string]int { // key 是无序的, 迭代的时候, 不保证顺序
		"S": 100,
		"A": 90,
		"B": 80,
		"C": 70,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	isS, hehe := m["S"] // map 的取值, 如果没有这个 key 会返回一个空值, 后面跟第二个变量, 可以用作判断是否有这个 key
	fmt.Println("map的key S:",isS, hehe)
}

func mapHandle () {
	og := map[string]string {
		"A": "AVal",
		"B": "BVal",
		"C": "CVal",
	}
	fmt.Println("初始map:", og)
	delete(og, "A") // 删除
	fmt.Println("删除操作后的mao", og)
}

func main () {
	//createMap()
	//iteration()
	mapHandle()
}
