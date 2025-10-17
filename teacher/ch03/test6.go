package main

import (
	"fmt"
	"time"
)

func main() {
	// 一个没有条件的 for 循环就是无限循环
	for {
		fmt.Println("我是一个无限循环，1秒后会再次打印...")
		time.Sleep(1 * time.Second)
		// 这种循环通常需要一个内部的 break 条件来退出
	}
}
