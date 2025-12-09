package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

/*
任务一：动态购物清单 (使用 fmt.Scanln)
要求：
1. 从键盘读取输入，直到用户输入 "done"。
2. 使用 fmt.Scanln()，注意它只读取到第一个空格。
3. 忽略空回车（处理 "unexpected newline" 错误）。
4. 将每次的输入（非 "done"）存入一个 string 切片。
5. 打印出完整的清单和完成时间。
*/
func main() {
	fmt.Println("--- 任务一：动态购物清单 ---")
	fmt.Println("请输入物品（*不含空格*），输入 'done' 结束:")
	// TODO 1: 声明一个空的 string 切片，变量名为 shoppingList 用于存放物品
	var shoppingList []string
	// TODO 2: 使用无限循环 for {}
	for {
		fmt.Print("> ") // 打印提示符
		var item string
		// TODO 3: 使用 fmt.Scanln(&item) 读取一个单词
		_, err := fmt.Scanln(&item)
		// TODO 4: 错误处理
		if err != nil {
			if err == io.EOF {
				// 用户按下了 Ctrl+D，视为输入结束
				break
			}
			// 如果用户只按了回车，Scanln 会返回 "unexpected newline" 错误
			// 我们捕获这个错误并继续循环，等待下一次输入
			if err.Error() == "unexpected newline" {
				fmt.Println("物品不能包含空格，请重新输入")
				continue
			}
			// 其他未知错误，打印并退出
			fmt.Println("读取输入时出错:", err)
			break
		}
		// TODO 5: 清理字符串并检查是否为 "done"
		item = strings.TrimSpace(item)
		if item == "done" {
			break
		}
		// TODO 6: 如果文本不为空，则使用 append() 将其添加到 shoppingList
		if item != "" {
			shoppingList = append(shoppingList, item)
		}
	}
	fmt.Println("\n--- 您的购物清单 ---")
	// TODO 7: 使用 for...range 遍历 shoppingList，并打印每个物品（例如："1: 苹果"）
	for i, item := range shoppingList {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	// TODO 8: 获取当前时间，并使用 "2006-01-02 15:04:05" 格式打印
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("清单完成时间：%s\n", now)
}

/*
-------------
预期输出 (示例):
-------------
--- 任务一：动态购物清单 ---
请输入物品（*不含空格*），输入 'done' 结束:
> apple
> banana
> (用户在这里只按了回车)
> milk
> (用户输入了 'ice cream'，Scanln 只读取了 'ice')
> (Scanln 丢弃了 'cream' 并返回 'unexpected newline'，被我们捕获)
>
> done
--- 您的购物清单 ---
1: apple
2: banana
3: milk
4: ice
清单完成时间: 2025-11-17 09:00:00 (这里显示你运行时的实际时间)
*/
