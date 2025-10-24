package main

import "fmt"

// intSequence 函数返回了一个新的函数
// 这个返回的函数是一个闭包
func intSequence() func() int {
	i := 0 // 变量 i 被闭包引用
	// 返回一个匿名函数
	return func() int {
		i++ // 每次调用时，i 都会递增
		return i
	}
}
func main() {
	// nextInt 是一个函数，它“绑定”了它自己的那个 i
	nextInt := intSequence()
	fmt.Println(nextInt()) // 输出 1
	fmt.Println(nextInt()) // 输出 2
	fmt.Println(nextInt()) // 输出 3
	// 创建一个新的闭包，它有自己独立的 i
	newInts := intSequence()
	fmt.Println(newInts()) // 输出 1
	fmt.Println(nextInt()) // 输出 ?
}
