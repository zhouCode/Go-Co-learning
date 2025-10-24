package main

import "fmt"

// 演示命名返回值
// (result int, remainder int) 这就是命名返回值
func divideWithRemainder(a, b int) (result int, remainder int) {
	// result 和 remainder 在这里已经可用，它们是 int 零值 (即 0)
	if b == 0 {
		return 0, 0 // 即使使用了命名返回，也可以显式返回
	}
	result = a / b
	remainder = a % b

	// 裸露的 return 会自动返回 result 和 remainder 的当前值
	return
}

func main() {
	res, rem := divideWithRemainder(10, 3)
	// 输出：10 / 3 = 3 ... 1
	fmt.Printf("10 / 3 = %d ... %d\n", res, rem)
}
