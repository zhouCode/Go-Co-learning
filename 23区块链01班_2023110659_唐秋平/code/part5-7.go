package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("猜数字游戏开始")
	randNum := rand.Intn(100)
	var guess int
	fmt.Println("请输入你猜的数字:")
	for {
		fmt.Scanln(&guess)
		fmt.Sprintf("你猜的数字f是:%d", guess)
		if guess == randNum {
			fmt.Println("恭喜你猜对了!")
			break
		} else if guess < randNum {
			fmt.Println("你猜的数字太小了")
		} else {
			fmt.Println("你猜的数字太大了")
		}
	}
}
