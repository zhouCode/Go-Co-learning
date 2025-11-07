package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1.Parse 类函数
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.ParseInt("0", 10, 64))
	fmt.Println(strconv.ParseInt("123x", 10, 64))

	//2.Format 类函数
	fmt.Printf("%T \n", strconv.FormatInt(123, 10))
	fmt.Printf("%T", strconv.FormatInt(0, 10))
}
