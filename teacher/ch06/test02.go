package main

import (
	"fmt"
	"strings"
)

// string包中的字符串处理函数
func main() {
	// 1. 检索字符串
	// Contains 检查字符串是否包含指定的子字符串,如果包含则返回true，否则返回false
	fmt.Println(strings.Contains("hello", "ll"))    // true
	fmt.Println(strings.Contains("hello", "world")) // false
	// 2. 分割字符串
	// Split 将字符串按照指定的分隔符进行分割，返回一个字符串切片
	fmt.Println(strings.Split("a,b,c", ",")) // [a b c]
	fmt.Println(strings.Split("a,b,c", ""))  // [a , b , c]
	fmt.Println(strings.Split("a,b,c", "x")) // [a,b,c]
	// 3. 大小写转换
	// ToLower 将字符串中的所有字符转换为小写
	fmt.Println(strings.ToLower("Hello")) // hello
	// ToUpper 将字符串中的所有字符转换为大写
	fmt.Println(strings.ToUpper("Hello")) // HELLO
	// 4. 修剪字符串
	// TrimSpace 修剪字符串首尾的空格
	fmt.Println(strings.TrimSpace("  hello  ")) // hello
	// 5. 比较字符串
	// EqualFold 不区分大小写地比较两个字符串是否相等
	fmt.Println(strings.EqualFold("hello", "HELLO")) // true

}
