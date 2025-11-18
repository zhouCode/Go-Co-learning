package main

import (
	"fmt"
	"time"
)

func main() {
	var items []string // 声明切片存储购物清单
	var input string

	for {
		fmt.Print("请输入购物清单物品（输入done结束）：")
		// 读取用户输入
		_, err := fmt.Scan(&input)
		if err != nil {
			// 忽略空输入（仅按回车）
			continue
		}

		// 输入"done"时结束循环
		if input == "done" {
			break
		}

		// 添加非空物品到切片
		if input != "" {
			items = append(items, input)
		}
	}

	// 打印购物清单
	fmt.Println("\n购物清单物品：")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// 打印当前时间（格式：年-月-日 时:分:秒）
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\n清单完成时间：%s\n", currentTime)
}
