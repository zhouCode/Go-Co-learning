//package main

//import "fmt"

// func main() {
// 	var a int = 10
// 	var b = 10
// 	c := 20 //第一次赋值才能用
// 	fmt.Println("hello world:", a, b, c)
// }

//package main
//import "fmt"
// func main() {
// 	a := 3
// 	switch a {
// 	case 1:
// 		fmt.Println("a=1")

// 	case 2:
// 		fmt.Println("a=2")
// 	default:
// 		fmt.Println("default")
// 	}
// }

//package main
//import "fmt"
// func main() {
//     fmt.Println("Hello, World!")
//     fmt.Println("欢迎 程蒙蒙 学习Go语言!")
//     fmt.Println("这是你的第一个Go程序")

//     // 基本变量声明
//     var name string = "程蒙蒙"
//     var age int = 20

//     fmt.Printf("学生姓名: %s\n", name)
//     fmt.Printf("年龄: %d\n", age)

//     // 简单的函数调用
//     greet(name)
// }

// // 问候函数
// func greet(name string) {
//     fmt.Printf("你好, %s! 欢迎来到Go语言的世界!\n", name)
// }

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Println("hello world")
// 	}
// }

package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("正在检查: i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				fmt.Println("找到了 跳出所有循环！")
				break OuterLoop
			}
		}
	}
}
