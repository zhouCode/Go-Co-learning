package main

import "fmt"

// func main() {
// 	var arr1 [3]int
// 	arr1[0] = 100
// 	fmt.Println("值：%v，类型：%T\n", arr1, arr1)
// 	//[100 0 0]
// 	arr2 := [...]int{1, 2, 3}
// 	fmt.Println("值：%v，类型：%T\n", arr2, arr2)
// 	//[1 2 3]
// 	var matrix [2][3]int
// 	matrix[0][0] = 1
// 	matrix[0][1] = 2
// 	matrix[0][2] = 3
// 	matrix[1][0] = 4
// 	matrix[1][1] = 5
// 	matrix[1][2] = 6
// 	fmt.Println(matrix)
// }
// func main() {
// 	a := [...]string{"US", "CHI", "JAp"}
// 	b := a
// 	fmt.Println("改前")
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	b[0] = "CNh"
// 	fmt.Println("改后")
// 	fmt.Println("a", a)
// 	fmt.Println("b", b)
// }
// func main() {
// 	s := make([]string, 3, 5)
// 	p := [...]int{1, 2, 3, 4, 5}
// 	c := p[1:4] //索引1-3

// 	fmt.Println(s)      //3个空字符串
// 	fmt.Println(len(c)) //3
// 	fmt.Println(cap(s)) //5
// 	fmt.Println(c)      //[2 3 4]
// }
func main() {
	s := make([]int, 1, 1)
	s = append(s, 10)
	fmt.Println("first", len(s), cap(s))
	s = append(s, 20)
	fmt.Println("second", len(s), cap(s))
	s = append(s, 30)
	fmt.Println("third", len(s), cap(s))
	s = append(s, 40, 59, 60)
	fmt.Println("fourth", len(s), cap(s))
}


