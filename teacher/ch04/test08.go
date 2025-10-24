package main

// 递归函数
import "fmt"

func factorial(n int) int {
	// 退出条件
	if n == 0 {
		return 1
	}
	// 递归调用
	return n * factorial(n-1)
}

func main() {
	fmt.Println("5! =", factorial(5)) // 输出 120
}
