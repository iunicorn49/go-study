package main

import "fmt"

func main () {
	str := "GO语言."

	for i, v := range str {
		fmt.Println(i, v, str[i])
	}
}
