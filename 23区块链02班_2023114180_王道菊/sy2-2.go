package main

import (
	"fmt"
	"strings"
)

func main() {
	// 给定包含多个单词的英文字符串（可替换为其他测试文本）
	text := "Go is a statically typed language Go is simple and efficient go is fun"

	// 分割字符串为单词切片（按空格分割）
	words := strings.Fields(text)

	// 用map存储单词频率（忽略大小写）
	freqMap := make(map[string]int)
	for _, word := range words {
		// 转换为小写统一处理
		lowerWord := strings.ToLower(word)
		freqMap[lowerWord]++
	}

	// 遍历map并打印结果
	fmt.Println("单词频率统计：")
	for word, count := range freqMap {
		fmt.Printf("%s: %d次\n", word, count)
	}
}
