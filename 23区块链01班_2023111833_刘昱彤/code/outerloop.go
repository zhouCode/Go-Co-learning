package main

import "fmt"

func main() {
outerloop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("正在检查： i=%d, j=%d", i, j)
			if i == 1 && j == 1 {
				fmt.Println("找到了，跳出所有循环")
				break outerloop
			}
		}
	}
}
