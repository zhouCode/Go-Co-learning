package main

import "fmt"

// 一个例子复习指针
func main() {
	var a int = 10
	// 1. 取地址 ……&a 表示取变量 a 的地址
	// 2. p 是一个指针变量，它存储了变量 a 的地址
	var p *int = &a
	fmt.Println(a)
	fmt.Println(p)
	// 3. *p 表示通过指针 p 访问到的变量 a 的值
	fmt.Println(*p)
}
