package main

// 作用域
import "fmt"

func main() {
	if v := 10; v > 5 {
		fmt.Println(v) // v 在这里可见
	} else {
		fmt.Println(v) // v 在这里也可见
	}
	// fmt.Println(v) // 在这里编译错误！v 的作用域结束了

	for i := 0; i < 3; i++ {
		fmt.Println(i) // i 在循环内可见
	}
	// fmt.Println(i) // 在这里编译错误！i 的作用域结束了
}
