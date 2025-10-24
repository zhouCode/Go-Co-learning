package main

import "fmt"

func main() {
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("匿名函数相乘结果：", multiply(5, 3))
	//定义并立即执行
	result := func(a, b int) int {
		return a * b
	}(5, 3)
	fmt.Println("立即执行结果：", result)
}
