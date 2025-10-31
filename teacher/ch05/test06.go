package main

import "fmt"

func main() {
	// 声明一个 map，但它是 nil
	// var m1 map[string]int

	// fmt.Println(m1 == nil) // 输出 true
	// fmt.Println(m1["key"]) // 读取 nil map 是安全的，返回 0

	// m1["key"] = 10 // 崩溃！panic: assignment to entry in nil map

	// --- 正确的方式 ---
	// 使用 make 来创建并初始化一个 map
	m2 := make(map[string]int)
	m2["key"] = 10
	fmt.Println(m2) // 输出 map[key:10]

	// 或者使用字面量初始化
	m3 := map[string]int{
		"alice": 25,
		"bob":   30,
	}
	fmt.Println(m3)
}
