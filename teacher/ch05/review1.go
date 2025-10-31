package main

import "fmt"

// 一个例子复习函数
func main() {
	a := 10
	b := 20
	c := add(a, b)
	fmt.Println(c)
	// 调用可变参数函数
	fmt.Println(sum(1, 2, 3, 4, 5))
	// 调用可变参数函数时，也可以直接传递一个数组
	var a1 = []int{1, 2, 3, 4, 5}
	fmt.Println(sum(a1...))
	// 函数变量
	f := add
	fmt.Println(f(10, 20))
	// 闭包
	
}

// 一个简单的函数，用于计算两个整数的和
func add(a int, b int) int {
	return a + b
}

// 可变参数
func sum(a ...int) int {
	var s int = 0
	for _, v := range a {
		s += v
	}
	return s
}
