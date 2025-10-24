package main

import "fmt"

// 匿名函数
func main() {
	// 1. 定义一个匿名函数并赋值给变量
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("multiply 现在指向的函数是: %T\n", multiply)
	fmt.Println("匿名函数相乘:", multiply(5, 3)) // 输出 15
	// 2. 定义并立即执行 (IIFE - Immediately Invoked Function Expression)
	// 注意最后面的 (5, 3) 是在调用它
	result := func(a, b int) int {
		return a * b
	}(5, 3)
	fmt.Printf("result 现在指向的函数是: %T\n", result)
	fmt.Println("立即执行的匿名函数:", result) // 输出 15
}
