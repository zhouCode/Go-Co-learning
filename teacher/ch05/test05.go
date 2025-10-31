package main

import "fmt"

func main() {
	var s []int // 声明一个 nil 切片 (len=0, cap=0)
	s = appendSlice(s)
	copySlice(s)
	deleteSlice(s)
}

func appendSlice(s []int) []int {
	fmt.Println("=======appendSlice 开始=======")
	// 第一次 append，s 是 nil，会分配新数组

	s = append(s, 10)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	// 第二次 append，容量可能够，也可能不够（取决于 Go 的策略）
	s = append(s, 20)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	s = append(s, 30)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	// 常见的错误写法：
	// append(s, 40) // 错误！这个 append 可能创建了新切片，但 s 没更新
	// 正确的写法：
	s = append(s, 40, 50, 60, 70)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Printf("=======appendSlice 结束=======\n")
	return s
}

func copySlice(s []int) {
	fmt.Println("=======copySlice 开始=======")
	// copy 操作
	s2 := make([]int, len(s))
	copy(s2, s) // 将 s 复制到 s2
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	fmt.Println("=======copySlice 结束=======")
}

func deleteSlice(s []int) {
	fmt.Println("=======deleteSlice 开始=======")
	// 删除第一个元素
	s = append(s[1:])
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	// 删除最后一个元素
	s = s[:len(s)-1]
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	// 删除中间元素
	a := int(len(s) / 2)
	s = append(s[:a], s[a+1:]...)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Println("=======deleteSlice 结束=======")
}
