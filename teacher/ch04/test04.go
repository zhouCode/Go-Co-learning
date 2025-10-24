package main

import "fmt"

// 函数变量
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	// 声明一个变量 f，它的类型是 "一个接收两个int并返回一个int的函数"
	// func(int, int) int
	var f func(int, int) int

	// 将函数 add 赋值给 f
	f = add
	fmt.Println("f(5, 3) =", f(5, 3)) // 输出 8

	// 将函数 subtract 赋值给 f
	f = subtract
	fmt.Println("f(5, 3) =", f(5, 3)) // 输出 2

	fmt.Printf("f 现在指向的函数是: %T\n", f)
}
