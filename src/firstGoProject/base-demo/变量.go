package main

import "fmt"

func emptyGroup() { // 声明空值
	var emptyNumber int
	var emptyString string
	fmt.Println("emptyNumber", emptyNumber) // 0
	fmt.Println("emptyString", emptyString) // ""
}

func main() {
	emptyGroup()
}
