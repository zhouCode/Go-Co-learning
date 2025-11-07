package main

import (
	"fmt"
	"strings"
)

func main() {
	//1.检索字符串
	fmt.Println(strings.Contains("hello", "ll"))
	fmt.Println(strings.Contains("hello", "world"))
	//2.分割字符串
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Split("a b c", " "))
	fmt.Println(strings.Split("a,b,c", "x"))
	//3.大小写转换
	fmt.Println(strings.ToLower("hello"))
	//转换成大写
	fmt.Println(strings.ToUpper("hello"))
}
