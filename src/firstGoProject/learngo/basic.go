package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ATOM",
		"age": "26",
		"hobby": "sleep",
	}
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)
}