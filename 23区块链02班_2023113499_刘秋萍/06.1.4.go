package main

import (
	"fmt"
	"time"
)

func main() {
	//1.时间解析
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:00:00")
	fmt.Println("时间解析:", t)
	//2.时间格式化
	now := time.Now()
	fmt.Println("当前时间:", now.Format("2006-01-02 15:04:05"))
	//3.时间计算
	fmt.Println("24小时后的时间：", now.Add(24*time.Hour))
	//Sub 时间差
	fmt.Println("时间差：", now.Sub(t))
	//4.时间比较
	fmt.Println("解析时间是否在当前时间之前：", t.Before(now))
	fmt.Println("解析时间是否在当前时间之后：", t.After(now))
	fmt.Println("时间比较：", now.Equal(t))

}
