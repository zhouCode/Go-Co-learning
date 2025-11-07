package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := "我爱go语言"
	//统计字节数
	for _, v := range s {
		fmt.Printf("c%", v)
	}
	fmt.Println()
	fmt.Println(len(s))
	//根据条件拆分
	fmt.Println(strings.Split("1 2 3", " "))
	//转小写
	fmt.Println(strings.ToLower("HELLO"))
	//转大写
	fmt.Println(strings.ToUpper("Hello"))
	//parse函数
	fmt.Println(strconv.ParseInt("123",10,64))
	//解析时间
	t, _ := time.Parse("2023-11-12 22:00:00","2006-01-02 15:04:05")
	fmt.Println("解析时间：",t)
	//时间格式化
	now := time.Now()
	fmt.Println("当前时间：",now.Format("2006-01-02 15:04:05"))
	//时间计算
	fmt.Println("100秒后：",now.Add(time.Second * 100).Format("2006-01-02 15:04:05"))
	fmt.Println("100分钟前：",now.Add(-time.Minute * 100).Format("2006-01-02 15:04:05"))
	//时间差
	fmt.Println("时间差：",now.Sub(t))
	
	//解析float
	fmt.Println(strconv.ParseFloat("3.14",64))
	//解析bool
	fmt.Println(strconv.ParseBool("true"))

	//随机种子
	seed := time.Now().UnixNano()
	fmt.Println(seed)
	//生成随机数
	fmt.Println("随机数：",rand.Intn(100))
	//生成随机浮点数
	fmt.Println("随机浮点数：",rand.Float64())
}
