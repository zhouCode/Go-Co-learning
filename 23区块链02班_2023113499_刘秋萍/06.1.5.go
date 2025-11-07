package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.随机数种子
	seed := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed)
	//2.生成随机数
	fmt.Println("随机数:", rand.Intn(100))
	fmt.Println("随机数:", r1.Intn(100))
	fmt.Println("随机浮点数:", r1.Float64())
}
