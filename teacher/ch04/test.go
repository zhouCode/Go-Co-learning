package main

import "fmt"

func main() {
	a := 10
	var ip *int = &a
	fmt.Printf("a 值的地址是: %x\n", &a)
	fmt.Printf("ip 现在指向的变量的值是: %v\n", *ip)
	// 1. 对于一个变量，想获取他的地址，就使用&[变量名]
	// 2. 对于一个指针变量，想获取他指向的变量，就使用*[指针变量名]
	// 3. 声明一个指针变量，使用 *[类型]
}
