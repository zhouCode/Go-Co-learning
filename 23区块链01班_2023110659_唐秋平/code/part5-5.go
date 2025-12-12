package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-11-06 12:34:56")
	fmt.Println("解析时间", t)
	now := time.Now()
	fmt.Println("当前时间:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("24小时后的时间:", now.Add(24*time.Hour))
	fmt.Println("时间差:", now.Sub(t))
	fmt.Println("解析时间是否在当前时间之前:", t.Before(now))
	fmt.Println("解析时间是否在当前时间之后:", t.After(now))
}
