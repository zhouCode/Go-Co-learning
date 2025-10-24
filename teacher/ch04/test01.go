package main

import (
	"errors"
	"fmt"
)

// 演示多返回值
// 1. 接收两个 int，返回一个 int (商) 和一个 error (错误)
// 2. 注意返回类型的声明：(int, error)
func divide(a, b int) (int, error) {
	if b == 0 {
		// 返回 0（作为 int 的默认值）和一个新的错误
		return 0, errors.New("除数不能为零")
	}
	// 返回计算结果和 nil (nil 在 Go 中表示 error 不存在)
	return a / b, nil
}

func main() {
	// 调用函数时，必须用两个变量来接收
	// result, err := divide(10, 2)
	// // 这就是 Go 处理错误的经典模式
	// if err != nil {
	// 	fmt.Println("发生错误:", err)
	// } else {
	// 	fmt.Println("10 / 2 =", result)
	// }

	// 再次尝试一个会出错的调用
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("发生错误:", err)
	// } else {
	// 	fmt.Println("10 / 0 =", result)
	// }

	// 如果你只关心其中一个返回值，可以用 _ 忽略另一个
	_, errOnly := divide(5, 0)
	fmt.Println("只关心错误:", errOnly)
}
