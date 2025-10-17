package main

import "fmt"

func main() {
	// 1. 初始化语句: i := 0
	// 2. 条件表达式: i < 5
	// 3. 后置语句:   i++
	for i := 0; i < 5; i++ {
		fmt.Println("当前 i 的值是:", i)
	}
}
