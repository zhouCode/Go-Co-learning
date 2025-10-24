package main

import "fmt"

func main() {
	A := func(a, b int) int {
		return a * b
	}
	fmt.Println("匿名函数相乘", A(5, 3))
}
