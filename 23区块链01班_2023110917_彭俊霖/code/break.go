package main

import "fmt"

func main(){
// 定义一个外层循环的标签
Outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++{
			fmt.Printf("正在检查：i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				fmt.Println("找到了！跳出所有循环！")
				break Outerloop
			}
		}
	}
	fmt.Println("循环结束")
}