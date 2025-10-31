package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// 连续运行两次，你可能会得到不同的输出顺序
	for k, v := range m {
		fmt.Printf("k: %s, v: %d | ", k, v)
	}
	fmt.Println()
	for k, v := range m {
		fmt.Printf("k: %s, v: %d | ", k, v)
	}
}
