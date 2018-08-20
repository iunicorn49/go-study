package main

import "fmt"

// 寻找最长不含有重复字符的子串
// abcabcbb => abc
// bbbbb => b
// pwwkew => wke

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, char := range []byte(s) { // 将字符转为8位
		if lastI, ok := lastOccurred[char]; ok && lastI >= start {
			start = lastOccurred[char] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[char] = i
	}
	return maxLength
}

func test(str string) {
	m := make(map[byte]int)
	for i, char := range []byte(str) {
		fmt.Println(i, char)
		m[char] = i
	}
	fmt.Println(m)
}

func main() {
	str := "pwwkew"
	test(str)
}
