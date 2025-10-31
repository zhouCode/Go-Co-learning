package main

import "fmt"

func main() {

	scores := map[string]int{
		"alice": 100,
		"bob":   85,
	}

	// 1. 检查 "alice"
	score, ok := scores["alice"]
	if ok {
		fmt.Println("Alice 的分数是:", score)
	} else {
		fmt.Println("Alice 不存在")
	}

	// 2. 检查 "charlie"
	score, ok = scores["charlie"]
	if ok {
		fmt.Println("Charlie 的分数是:", score)
	} else {
		fmt.Println("Charlie 不存在")
	}

	// 3. 如果你只关心值，不关心是否存在（不推荐）
	// bob_score := scores["bob"] // 85
	// charlie_score := scores["charlie"] // 0
	// 你无法区分 charlie 是真的0分，还是他根本不存在！
}
