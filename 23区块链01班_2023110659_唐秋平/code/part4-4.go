package main

import (
	"fmt"
)

func main() {
	var arr1 [3]int
	arr1[0] = 100
	fmt.Printf("值：%v,类型：%T、n", arr1, arr1)
	arr2 := [...]int{10, 20, 30, 40}
	fmt.Printf("值：%v,类型：%T、n", arr2, arr2)
}
