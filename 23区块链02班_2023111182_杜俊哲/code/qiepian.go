package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3, 5)
	parent := [...]int{10, 20, 30, 40, 50}
	child := parent[1:4]
	fmt.Println("slice:", s)
	fmt.Println("cap:", cap(s))
	fmt.Println("len:", len(s))
	fmt.Println(child)
	
	// 调用appendSlice函数
	var intSlice []int
	result := appendSlice(intSlice)
	fmt.Println("最终结果:", result)
}
func appendSlice(s []int) []int {
	s = make([]int, 1, 1)
	s = append(s, 10)
	fmt.Printf("s:%v,cap:%d,len:%d\n", s, cap(s), len(s))
	s = append(s, 20)
	fmt.Printf("s:%v,cap:%d,len:%d\n", s, cap(s), len(s))
	s = append(s, 30)
	fmt.Printf("s:%v,cap:%d,len:%d\n", s, cap(s), len(s))
	s = append(s, 40, 50, 60)
	fmt.Printf("s:%v,cap:%d,len:%d\n", s, cap(s), len(s))
	return s
}
