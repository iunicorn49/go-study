package main

import "fmt"

func main() {
	n := 3

	switch {
	case n > 1:
		fmt.Println("n > 1")
	case n > 2:
		fmt.Println("n > 2")
	default:
		fmt.Println("OK")
	}

}