package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := "我爱Go语言"
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	fmt.Println(len(s))
	for i, ch := range []byte(s) {
		fmt.Printf("第%d个字符为：%c\n", i, ch)
	}
	fmt.Println()
}
func main1() {
	//1.检索字符串
	fmt.Println(strings.Contains("hello", "ll"))    //true
	fmt.Println(strings.Contains("hello", "world")) //false
	//2.分割字符串
	fmt.Println(strings.Split("a,b,c", ",")) //[a b c]
	fmt.Println(strings.Split("a,b,c", " ")) //[a , b , c]
	fmt.Println(strings.Split("a,b,c", "x")) //[a,b,c]
	//3.大小写转换
	fmt.Println(strings.ToUpper("hello")) //HELLO
	fmt.Println(strings.ToLower("HELLO")) //hello
	//4.修剪字符串
	fmt.Println(strings.TrimSpace("  hello world  ")) //hello world
	//5.比较字符串

}

func main2() {
	//1.Parse类函数
	fmt.Println(strconv.ParseInt("123", 10, 64)) //123
	fmt.Println(strconv.ParseInt("0", 10, 64))   //0
	fmt.Println(strconv.ParseInt("123x", 10, 64))
	//2.Format类函数
	fmt.Println("%T\n", strconv.FormatInt(123, 10)) //123
	fmt.Println("%T", strconv.FormatInt(0, 10))     //0
}
func main3() {
	//1.正则表达式的匹配
	fmt.Println(regexp.MustCompile(`\d+`).MatchString("123"))  //true
	fmt.Println(regexp.MustCompile(`\d+`).MatchString("a123")) //false
	//2.正则表达式的替换
	fmt.Println(regexp.MustCompile(`\d+`).ReplaceAllString("a123b456", "X")) //aXbX
	//3.正则表达式的提取
	fmt.Println(regexp.MustCompile(`\d+`).FindString("a123b456")) //123
}
func main4() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-11-07 12:34:56")
	//1.解析时间字符串
	fmt.Println("解析时间", t) //2023-11-07 12:34:56 +0000 UTC
	//2.格式化时间
	now := time.Now()
	fmt.Println("当前时间", now.Format("2006-01-02 15:04:05")) //2023-11-07 12:34:56
	//3.时间的计算
	fmt.Println("24小时后的时间", now.Add(24*time.Hour)) //2023-11-08 12:34:56
	//sub 计算时间差
	fmt.Println("时间差", now.Sub(t)) //24h0m0s
	//4.时间的比较
	fmt.Println("时间比较", now.After(t))  //true
	fmt.Println("时间比较", now.Before(t)) //false
}
func main5() {
	//1.数学函数
	fmt.Println("绝对值", math.Abs(-123)) //123
	fmt.Println("平方根", math.Sqrt(9))   //3
	fmt.Println("立方根", math.Cbrt(27))  //3
}
func main6() {
	//1.随机数
	fmt.Println("随机数", rand.Intn(100)) //0-99的随机整数
	fmt.Println("随机数", rand.Float64()) //0.0-0.9999999999的随机浮点数
	//2.随机数的种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println("随机数", rand.Intn(100)) //0-99的随机整数
}
func main7() {
	fmt.Println("猜数字游戏开始")
	//1.生成随机数
	randNum := rand.Intn(100) + 1
	//2.猜数字游戏
	var guessNum int
	fmt.Println("请输入你猜的数字(1-100)：")
	fmt.Scanln(&guessNum)
	for guessNum != randNum {
		if guessNum < randNum {
			fmt.Println("你猜的数字太小了")
		} else {
			fmt.Println("你猜的数字太大了")
		}
		fmt.Println("请重新输入你猜的数字(1-100)：")
		fmt.Scanln(&guessNum)
	}
	fmt.Println("恭喜你，猜对了！")

}
