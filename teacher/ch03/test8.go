package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // 当 i=3 时，跳过下面的 Println，直接进入下一次循环 (i=4)
		}
		if i == 7 {
			break // 当 i=7 时，直接终止并跳出整个 for 循环
		}
		fmt.Println(i)
	}
	// 输出: 0, 1, 2, 4, 5, 6
}
