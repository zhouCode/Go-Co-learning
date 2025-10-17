package main

import "fmt"

func main() {
	// 定义一个外层循环的标签
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("正在检查: i:%d, j:%d\n", i, j)
			if i == 1 && j == 1 {
				fmt.Println("找到目标，跳出循环！")
				break OuterLoop // 使用 break + 标签，直接跳出到 OuterLoop 那一层
			}
		}
	}
	fmt.Println("进程结束")
}
