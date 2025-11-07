package main

import "fmt"

func main() {
	s := "我爱Go语言"
	//1.遍历字符串
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	//2.字符串长度
	fmt.Println(len(s))
	//3.遍历所有字符串
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x ", i, ch)
	}
	fmt.Println()
}
