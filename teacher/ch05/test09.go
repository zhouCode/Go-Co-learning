package main

import "fmt"

func main() {
	// delete 函数
	scores := map[string]int{
		"alice":   100,
		"bob":     85,
		"charlie": 90,
	}

	// 删除 "alice"
	delete(scores, "alice")
	fmt.Println(scores) // 输出 map[bob:85]

	// 清空map中所有元素
	for key := range scores {
		delete(scores, key)
	}
	fmt.Println(scores) // 输出 map[]
	// 更好的写法
	scores = make(map[string]int)
	fmt.Println(scores) // 输出 map[]

	// map是引用类型
	// 当你将一个 map 赋值给另一个变量时，
	// 两个变量都引用了同一个底层数据结构
	// 对一个变量的操作会影响到另一个变量
	scores2 := scores
	delete(scores2, "bob")
	fmt.Println(scores) // 输出 map[charlie:90]
}
