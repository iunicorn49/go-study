package main

import (
	"fmt"
	"unicode/utf8"
)

func printS (s string) { // UTF-8 中文占3字节, 英文为1字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch 是一个 rune 一个4字节的整数
		fmt.Printf("(%d %X) ", i, ch) // 留意 i 的变化, i 对应的是字节, 碰到中文会跳位
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 字符长度为9
	bytes := []byte(s)
	for len(bytes) > 0 { // 一个字符一个字符的打印
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) { // 同上, rune 是字符型单位, i 是连续的
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
// 字符串处理 总结
// 使用 range 遍历 string 的时候, i 对于的是字节, 一个中文占3字节, 所以, i 不连续
// 使用 utf8.RuneCountInString 可以获取字符数量(无关中英文)
// 使用 len 获得的是字节长度, 会受到中文的影响
// 使用 []byte 获得字节

func lengthOfNonRepeatingSubStrNoReturn (s string) {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, char := range []rune(s) { // 将字符转为8位
		if lastI, ok := lastOccurred[char]; ok && lastI >= start {
			start = lastOccurred[char] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[char] = i
	}
	fmt.Println("maxLength:", maxLength)
}

func main () {
	//s := "Yes我爱慕课网!"
	//printS(s)
	// 寻找最长不含有重复字符的子串
	lengthOfNonRepeatingSubStrNoReturn("呵呵哒") // 2
}