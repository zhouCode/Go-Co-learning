package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 1:
		fmt.Println("a等于1")
	case 2:
		fmt.Println("a等于2")
	default:
		fmt.Println("a不等于1也不等于2")
	}
}
