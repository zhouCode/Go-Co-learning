package main

import (
	"fmt"
	"strconv"
)

// strconv包中常用函数
func main() {
	// 1. Parse类函数
	// ParseBool 将字符串解析为 int 值
	fmt.Println(strconv.ParseInt("123", 10, 64))  // 123, nil
	fmt.Println(strconv.ParseInt("0", 10, 64))    // 0, nil
	fmt.Println(strconv.ParseInt("123x", 10, 64)) // 0, strconv.ErrSyntax
	// 2. Format类函数
	// FormatBool 将 int 值格式化为字符串
	fmt.Printf("%T \n", strconv.FormatInt(123, 10)) // 123 123
	fmt.Printf("%T", strconv.FormatInt(0, 10))      // 0 0
}
