package main

import "fmt"

// 在 C++/Java 中，这是合法的
// if (score > 60) printf("及格了");

// 在 Go 中，下面的写法是错误的！
//
// if score > 60
//     fmt.Println("及格了") // 编译错误！
//

func main() {
	score := 80
	// 正确的 Go 写法
	if score > 60 {
		fmt.Println("及格了")
	}
}
