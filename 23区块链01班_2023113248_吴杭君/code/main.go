package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("欢迎 吴杭君 学习Go语言!")
    fmt.Println("这是你的第一个Go程序")
    
    // 基本变量声明
    var name string = "吴杭君"
    var age int = 20
    
    fmt.Printf("学生姓名: %s\n", name)
    fmt.Printf("年龄: %d\n", age)
    
    // 简单的函数调用
    greet(name)
}

// 问候函数
func greet(name string) {
    fmt.Printf("你好, %s! 欢迎来到Go语言的世界!\n", name)
}
