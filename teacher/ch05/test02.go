package main

import "fmt"

// Go语言中，数组是“值类型”
// 这个函数接收的是 arr 的一个完整副本
func modifyArray(arr [3]int) {
	arr[0] = 999 // 只会修改这个副本
	fmt.Println("In function:", arr)
}

func main() {
	// 声明并初始化
	myArray := [3]int{10, 20, 30}

	fmt.Println("Before call:", myArray)

	modifyArray(myArray) // 传递的是 myArray 的一个拷贝

	fmt.Println("After call:", myArray) // myArray 自身没有被改变
}
