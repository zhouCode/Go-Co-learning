package main

import "fmt"

func main() {
	a := 10
	var ip *int = &a
	fmt.Println("现有指针a的地址为：%x\n", &a)
	//对于一个变量，想获取他的地址，就使用&[变量名]
	//对于一个指针变量，想获取他指向的变量的值，就使用*[指针变量名]
	//声明一个指针变量，使用*[类型]
	fmt.Println("指针p的地址为：%x\n", &ip)
	fmt.Println("指针p指向的变量的值为：%d\n", *ip)
}
