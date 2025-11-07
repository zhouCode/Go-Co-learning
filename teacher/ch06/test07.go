package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数的核心方法
func main() {
	// 1. 随机数种子
	// NewSource 使用指定的种子值返回一个新的伪随机数生成器
	seed := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed)
	// 2. 生成随机数
	fmt.Println("随机数: ", rand.Intn(100)) // 0-99 之间的随机数
	fmt.Println("随机数: ", r1.Intn(100))   // 0-99 之间的随机数
	fmt.Println("随机浮点数: ", r1.Float64()) // 0.0-0.9999999999999999 之间的随机浮点数
}
