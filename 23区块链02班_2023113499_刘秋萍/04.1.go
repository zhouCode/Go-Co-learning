package main

import "fmt"

func main() {
	multipart := func(a, b int) int {
		return a * b
	}
	fmt.Printf("multipart现在指向的函数是：%T\n", multipart)
	fmt.Println("匿名函数相乘：", multipart(5, 3)) //输出15
	result := func(a, b int) int {
		return a * b
	}(5, 3)
	fmt.Println("匿名函数相乘结果：", result)
}
