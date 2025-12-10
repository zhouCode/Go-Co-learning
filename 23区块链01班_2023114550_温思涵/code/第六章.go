package main

import (
	"fmt"
	"time"
)

func main() {
	// s := "我爱Go语言"
	// //range遍历字符串
	// for _, value := range s {
	// 	fmt.Printf("%c", value)
	// }
	// fmt.Println()
	// //字符串的长度
	// fmt.Printf("字符串的长度：%d", len(s))
	// //遍历所有字节
	// for i, ch := range []byte(s) {
	// 	fmt.Printf("%d:%x", i, ch)
	// }
	// fmt.Println()
	// fmt.Println(strconv.ParseInt("123", 10, 64))
	// fmt.Println(strconv.ParseInt("0", 10, 64))
	// fmt.Println(strconv.ParseInt("123x", 10, 64))
	//1.时间解析（Parse将字符串解析为时间）
	t, _ := time.Parse("2006-01-02 15:04:05", "2025-11-07 15:00:00")
	fmt.Println("解析时间：", t)
	//2.时间格式化（Format将时间格式化为字符串）
	now := time.Now()
	fmt.Println("当前时间：", now.Format("2006-01-02 15:04:05"))
	//3.时间差（Sub返回两个时间的差值）
	fmt.Println("时间差：", now.Sub(t))
	//4.时间比较（Before、After、Equal）
	fmt.Println("解析时间是否在当前时间之前：", t.Before(now))
	fmt.Println("解析时间是否在当前时间之后：", t.After(now))
	fmt.Println("解析时间是否等于当前时间：", t.Equal(now))
}
