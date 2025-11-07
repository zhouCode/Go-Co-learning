package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomInt := rand.Intn(101)
	fmt.Printf("随机整数 (0-100): %d\n", randomInt)
	randomInt2 := rand.Intn(10) + 1
	fmt.Printf("随机整数 (1-10): %d\n", randomInt2)

	randomFloat := rand.Float64()
	fmt.Printf("随机浮点数 (0.0-1.0): %.6f\n", randomFloat)

	randomFloat2 := rand.Float64() * 100
	fmt.Printf("随机浮点数 (0.0-100.0): %.2f\n", randomFloat2)

	min := 10.5
	max := 20.5
	randomFloat3 := min + rand.Float64()*(max-min)
	fmt.Printf("随机浮点数 (%.1f-%.1f): %.2f\n", min, max, randomFloat3)

	fmt.Println("\n=== 生成多个随机数 ===")

	fmt.Print("5个随机整数: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", rand.Intn(100))
	}
	fmt.Println()
	fmt.Print("5个随机浮点数: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%.2f ", rand.Float64()*10)
	}
	fmt.Println()
}
