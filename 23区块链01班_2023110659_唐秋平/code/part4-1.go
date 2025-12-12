package main

import "fmt"

func main() {
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("multiply 现在指向的函数是：%T、n", multiply)
	fmt.Println("匿名函数相乘：", multiply(5, 3))
}
