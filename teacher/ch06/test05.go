package main

import (
	"fmt"
	"time"
)

// time 包的核心方法
func main() {
	// 1. 时间解析	// Parse 将字符串解析为时间
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-08-01 12:34:56")
	fmt.Println("解析时间: ", t) // 2023-08-01 12:34:56 +0000 UTC
	// 2. 时间格式化	// Format 将时间格式化为指定的字符串
	now := time.Now()
	fmt.Println("当前时间: ", now.Format("2006/01/02 15:04:05")) // 2025-11-07 15:00:00
	// 3. 时间计算	// Add 将时间加上指定的时间间隔
	fmt.Println("24小时后的时间: ", now.Add(24*time.Hour)) // 2025-11-08 15:00:00
	// Sub 计算时间差
	fmt.Println("时间差: ", now.Sub(t)) // 19874h18m47.8785712s
	// 4. 时间比较
	// Before 判断时间是否在指定时间之前
	fmt.Println("解析时间是否在当前时间之前: ", t.Before(now)) // true
	// After 判断时间是否在指定时间之后
	fmt.Println("解析时间是否在当前时间之后: ", t.After(now)) // false
}
