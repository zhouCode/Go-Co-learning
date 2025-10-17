package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
		fallthrough //实现穿透效果
	case 2:
		fmt.Println("2")
		fallthrough //实现穿透效果
	default:
		fmt.Println("default")
	}
}
