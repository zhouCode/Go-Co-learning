package main

import "fmt"

func main() {
	sum := 1
	// 只有一个条件判断，这就是 Go 版本的 while 循环
	for sum < 100 {
		sum += sum // 每次翻倍
	}
	fmt.Println("最终 sum 的值是:", sum) // 输出 128
}
