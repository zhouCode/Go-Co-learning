package code

import "fmt"

func main() {
	// 清晰的输出标识
	fmt.Println("=== 开始执行条件判断程序 ===")
	
	score := 85
	fmt.Printf("当前分数: %d\n", score)
	
	if score >= 60 {
		fmt.Println("结果: 及格 ✓")
	} else {
		fmt.Println("结果: 不及格 ✗")
	}
	if score2 := 70; score2 >= 60 {
		fmt.Println("结果: 及格 ✓")
	} else {
		fmt.Println("结果: 不及格 ✗")
	}
	}
