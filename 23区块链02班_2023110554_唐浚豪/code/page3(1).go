package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("欢迎学习Go语言！")
    
    // 简单的条件判断
    score := 85
    if score >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }
    
    // if内初始化语句
    if score2 := 85; score2 >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }
}

