package main

import (
	"fmt"
	"math"
)

// math 包核心方法
func main() {
	// 最大值
	fmt.Println("最大值: ", math.Max(2, 3)) // 3
	// 最小值
	fmt.Println("最小值: ", math.Min(2, 3)) // 2
	// 绝对值
	fmt.Println("绝对值: ", math.Abs(-2.5)) // 2.5
	// Pow 幂运算
	fmt.Println("2的3次方: ", math.Pow(2, 3)) // 8
	// Sqrt 平方根
	fmt.Println("4的平方根: ", math.Sqrt(4)) // 2
	// Round 四舍五入到最近的整数
	fmt.Println("3.5四舍五入: ", math.Round(3.5)) // 4
	// Sin 正弦函数
	fmt.Println("30度的正弦值: ", math.Sin(30*math.Pi/180)) // 0.5
	// Log 自然对数
	fmt.Println("2的自然对数: ", math.Log(2)) // 0.6931471805599453
}
