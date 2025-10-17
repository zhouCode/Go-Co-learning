package main

import "fmt"

//定义一个外层循环的标签
func main() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("正在检查：i=%d,j=%d\n", i, j)
			if i == 1 && j == 1 {
				fmt.Println("发现异常：i=%d,j=%d\n", i, j)
				break OuterLoop //当i等于1,j等于1时，跳出外层循环
			}
		}
	}
	fmt.Println("检查完成")
}
