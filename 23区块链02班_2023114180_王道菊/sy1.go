package main

import "fmt"

/*
任务一：实现一个多返回值的奇偶性判断函数
要求：
1. 接收一个整数 n。
2. 返回两个值：(bool, string)
3. 如果 n 是偶数，返回 (true, "偶数")
4. 如果 n 是奇数，返回 (false, "奇数")
*/
// TODO: 完成函数签名和函数体
func checkOddEven(n int) (bool, string) {
	// TODO: 使用 if-else 结构和 % 运算符实现判断逻辑
	if n%2 == 0 {
		return true, "偶数"
	} else {
		return false, "奇数"
	}
}

/*
任务二：实现一个状态计数器（闭包）
要求：
1. counter() 函数返回一个 func() int 类型的函数（一个闭包）。
2. 这个闭包捕获一个外部变量（例如 i）。
3. 每次调用这个闭包时，它应该使 i 递增，并返回 i 的当前值（从 1 开始）。
*/func counter() func() int {
	i := 0 // TODO: 在这里定义一个变量，它将被下面的闭包捕获
	return func() int {
		i++      // TODO: 在这里实现递增逻辑并返回
		return i // 这是一个占位符，请替换它
	}
}

/*
任务三：实现一个斐波那契数列生成器（闭包）
要求：
1. fibonacci() 函数返回一个 func() int 类型的函数（一个闭包）。
2. 这个闭包捕获两个状态变量（例如 a 和 b）来跟踪当前的斐波那契数。
3. 每次调用闭包时，它返回序列中的下一个数字。
   (序列: 0, 1, 1, 2, 3, 5, ...)
*/
func fibonacci() func() int {
	a, b := 0, 1 // TODO: 在这里定义并初始化状态变量
	return func() int {
		// TODO: 实现返回当前斐波那契数，并更新状态以便下次调用的逻辑
		// 提示：你需要一个变量来保存当前要返回的值，
		// 然后再更新 a 和 b 的状态。
		temp := a
		a = b
		b = temp + b
		return temp // 这是一个占位符，请替换它
	}
}
func main() {
	fmt.Println("--- 任务一：奇偶性判断 ---")
	isEven1, label1 := checkOddEven(4)
	fmt.Printf("数字 4: 是偶数吗? %v, 标签: %s\n", isEven1, label1) // 预期: true, 偶数

	isEven2, label2 := checkOddEven(7)
	fmt.Printf("数字 7: 是偶数吗? %v, 标签: %s\n", isEven2, label2) // 预期: false, 奇数
	fmt.Println("\n--- 任务二：计数器 ---")
	// c1 是一个闭包，它拥有自己的状态 i
	c1 := counter()
	fmt.Println(c1()) // 预期: 1
	fmt.Println(c1()) // 预期: 2
	fmt.Println(c1()) // 预期: 3
	fmt.Println("--- 独立的计数器 ---")
	// c2 是另一个闭包，它拥有自己独立的的状态 i
	c2 := counter()
	fmt.Println(c2()) // 预期: 1
	fmt.Println(c1()) // 预期: 4 (验证 c1 不受 c2 影响)
	fmt.Println("\n--- 任务三：斐波那契数列 ---")
	f := fibonacci()
	fmt.Println("斐波那契数列前 10 项:")
	// TODO: 循环调用
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	// 预期: 0 1 1 2 3 5 8 13 21 34
	fmt.Println()
}

/*
-------------
预期输出:
-------------
--- 任务一：奇偶性判断 ---
数字 4: 是偶数吗? true, 标签: 偶数
数字 7: 是偶数吗? false, 标签: 奇数
--- 任务二：计数器 ---
1
2
3
--- 独立的计数器 ---
1
4
--- 任务三：斐波那契数列 ---
斐波那契数列前 10 项:
0 1 1 2 3 5 8 13 21 34
*/
