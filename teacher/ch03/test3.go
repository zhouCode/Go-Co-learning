package main

import (
	"errors"
	"fmt"
)

// 检查密码长度，如果小于8位就返回一个错误
func checkPassword(password string) (int, error) {
	if len(password) < 8 {
		return 0, errors.New("密码长度不能小于8位")
	}
	return len(password), nil
}

func main() {
	// 这就是 "if with a short statement"
	// 1. 调用 checkPassword 并声明 length 和 err 变量
	// 2. 紧接着用分号 ; 隔开，进行条件判断 err != nil
	// 3. 变量 length 和 err 的作用域只在 if/else if/else 结构内
	if length, err := checkPassword("123456"); err != nil {
		// 如果 err 不是 nil，说明出错了
		fmt.Println("检查密码时发生错误:", err)
		// 在这里无法访问 length，因为它是在成功路径上才有效的值
	} else {
		// 如果 err 是 nil，说明成功了
		fmt.Println("密码合法，长度为:", length)
	}

	// 在 if 结构外面尝试访问 err 或 length，会编译错误！
	// fmt.Println(err) // uncommenting this line will cause a compile error
}
