package main
import (
"fmt"
"strings"
)
/*
任务二：单词频率统计器
要求：
1. 给定一个字符串 text。
2. 统计其中每个单词出现的次数（忽略大小写）。
3. 使用 map 存储结果。
4. 打印出每个单词及其频率。
*/
func main() {
	fmt.Println("--- 任务二：单词频率统计器 ---")
	text := "Go is a great language Go is fun and Go is powerful go"
	fmt.Println("原始文本:", text)
	// TODO 1: 使用 strings.ToLower() 将 text 转换为全小写
	lowerText := strings.ToLower(text)
	// TODO 2: 使用 strings.Fields() 将 lowerText 分割成单词切片
	words := strings.Fields(lowerText)
	// TODO 3: 使用 make() 创建一个 map[string]int 来存储单词计数
	wordCounts := make(map[string]int)
	// TODO 4: 遍历 words 切片
	// 在循环中，更新 map 中的计数 (wordCounts[word]++)
	for _, word := range words {
		wordCounts[word]++
	}
	fmt.Println("--- 单词频率 ---")
	// TODO 5: 遍历 wordCounts map，打印每个单词和它的数量
	// 格式："<word>": <count>
	for word, count := range wordCounts {
		fmt.Printf("%q: %d\n", word, count)
	}
}
/*
-------------
预期输出:
-------------
--- 任务二：单词频率统计器 ---
原始文本: Go is a great language Go is fun and Go is powerful go
--- 单词频率 ---
'go': 4
'is': 3
'a': 1
'great': 1
'language': 1
'fun': 1
'and': 1
'powerful': 1
(Map 的输出顺序可能是随机的)
*/