package main

import "fmt"

func intSquare() func() int {
	i := 0 //变量i被闭包引用
	//返回一个匿名函数
	return func() int {
		i++ //每次调用是，i都会递增
		return i
	}
}
func main() {
	//nextInt 是一个闭包函数，它引用了外部变量 i
	nextInt := intSquare()
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	//建立一个新的闭包函数，它引用了外部变量 i 的新值
	nextInts := intSquare()
	fmt.Println(nextInts()) //1
	fmt.Println(nextInt())
}
