package main

import "fmt"

func main() {
	// 底层数组是 [10, 20, 30, 40, 50]
	parent := []int{10, 20, 30, 40, 50}

	// child 是 [20, 30, 40]
	child := parent[1:4] // 索引 1 到 3 (不含 4)

	fmt.Println("Parent:", parent)
	fmt.Println("Child:", child)

	// *** 重点来了 ***
	// 修改 child 的元素
	child[0] = 99 // 对应 parent[1]

	// 查看 parent
	fmt.Println("After modify child, Parent is:", parent) // 输出 [10 99 30 40 50]
}
