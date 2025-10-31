package main

import "fmt"

func main() {
	// 创建一个 []string 类型的切片
	// 长度为 3，容量为 5
	// 这会分配一个能容纳 5 个 string 的底层数组
	s := make([]string, 3, 5)
	// 底层数组是 [10, 20, 30, 40, 50]
	parent := []int{10, 20, 30, 40, 50}
	// child 是 [20, 30, 40]
	child := parent[1:4]         // 索引 1 到 3 (不含 4)
	fmt.Println("Slice: ", s)    // 输出 [  ] (3个空字符串)
	fmt.Println("Len: ", len(s)) // 输出 3
	fmt.Println("Cap: ", cap(s)) // 输出 5
	fmt.Println("child: ", child)
	// 你也可以用字面量创建（常用）
	// 这会自动创建一个底层数组，并让切片指向它
	s_literal := []int{10, 20, 30}
	fmt.Println(s_literal, len(s_literal), cap(s_literal)) // 输出 [10 20 30] 3 3
}
