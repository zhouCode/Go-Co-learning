package main

import "fmt"

func main() {
	// 在main函数内部定义匿名函数并赋值给变量
	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Println("Hello, World!")
	fmt.Println(multiply(2, 3))
}
