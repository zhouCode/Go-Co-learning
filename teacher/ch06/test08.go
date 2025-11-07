package main

import (
	"fmt"
	"math/rand"
)

// 键盘输入 -- 猜数字游戏
func main() {
	// 键盘输入 -- 猜数字游戏
	fmt.Println("猜数字游戏开始")
	// 1. 生成随机数
	randNum := rand.Intn(100)
	// 2. 键盘输入
	var guess int
	// 3. 比较猜的数字和随机数
	fmt.Println("请输入你猜的数字: ")
	for {
		fmt.Scanln(&guess)
		fmt.Println("你猜的数字是: ", guess)
		if guess == randNum {
			fmt.Println("恭喜你猜对了")
			break
		} else {
			if guess > randNum {
				fmt.Println("你猜大了,请继续：")
			} else {
				fmt.Println("你猜小了,请继续：")
			}
		}
	}
}
