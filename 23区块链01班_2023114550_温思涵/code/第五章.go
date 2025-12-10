package main

import "fmt"

// func main() {
// 	var arr1 [3]int
// 	arr1[0] = 100

// 	fmt.Printf("值:%v,类型：%T\n", arr1, arr1)
// }

// func main() {
// 	a := [...]string{"USA", "China", "Japen"}
// 	b := a
// 	fmt.Println("修改前：")
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	b[0] = "Singapore"
// 	fmt.Println("修改后：")
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// }

// func main() {
// 	s := make([]string, 3, 5)
// 	parent := [...]int{10, 20, 30, 40, 50} //数组
// 	child := parent[1:4]                   //切片
// 	fmt.Println("Slice:", s)
// 	fmt.Println("Len:", len(s))
// 	fmt.Println("Cap:", cap(s))
// 	fmt.Println("Child:", child)
// }

func main() {
	s := make([]int, 1, 1)
	s = append(s, 10)
	fmt.Printf("s:%v,len:%d,cap:%d\n", s, len(s), cap(s))
	s = append(s, 20)
	fmt.Printf("s:%v,len:%d,cap:%d\n", s, len(s), cap(s))
	s = append(s, 30)
	fmt.Printf("s:%v,len:%d,cap:%d\n", s, len(s), cap(s))
	s = append(s, 40, 50, 60, 70)
	fmt.Printf("s:%v,len:%d,cap:%d\n", s, len(s), cap(s))
}
