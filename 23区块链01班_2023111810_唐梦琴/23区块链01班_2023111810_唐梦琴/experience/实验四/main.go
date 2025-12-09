package main

import (
	"fmt"
	"sync"
	"time"
)

/*
任务一：面包生产线 (Channel)
要求：
1. 创建一个 int 通道。
2. 启动一个协程（面包师）发送 5 个数字。
3. 主线程（顾客）接收并打印这些数字。
*/
func task1() {
	fmt.Println("\n--- 任务一：面包生产线 ---")
	// TODO 1: 创建一个 int 类型的通道 (channel)
	// breadBox := ...
	breadBox := make(chan int) // 这是一个占位符，请确认是否正确
	// 启动面包师协程
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("面包师: 正在制作面包 %d...\n", i)
			time.Sleep(50 * time.Millisecond) // 模拟制作时间
			// TODO 2: 将面包序号 i 发送到通道 breadBox 中
			// ...
			breadBox <- i
		}
		fmt.Println("面包师: 下班了 (关闭通道)")
		// TODO 3: 关闭通道，通知接收方没有更多数据了
		// ...
		close(breadBox)
	}()
	
	fmt.Println("顾客: 准备接收面包...")
	// TODO 4: 使用 for range 循环从 breadBox 通道中不断读取数据
	for bread := range breadBox {
		fmt.Printf("顾客: 收到面包 %d\n", bread)
	}
}

/*
任务二：并发打印机
要求：
1. 启动 3 个协程。
2. 每个协程打印 3 次消息。
3. 使用 sync.WaitGroup 确保主程序等待所有协程完成。
*/
func task2() {
	fmt.Println("--- 任务二：并发打印机 ---")
	// 声明 WaitGroup
	var wg sync.WaitGroup
	// 循环启动 3 个协程
	for i := 1; i <= 3; i++ {
		// TODO 1: 在启动协程前，通知 wg 增加一个任务计数
		wg.Add(1)
		// TODO 2: 使用 'go' 关键字启动一个匿名函数
		// 注意：我们将 i 作为参数 id 传递进去，避免闭包捕获循环变量的问题

		go func(id int) {
			// TODO 3: 确保函数结束时通知 wg 任务已完成
			// defer ...
			defer wg.Done()
			for count := 1; count <= 3; count++ {
				fmt.Printf("协程 %d: 第 %d 次打印\n", id, count)
				// 模拟耗时，帮助观察并发效果
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	fmt.Println("主线程: 等待所有协程结束...")
	// TODO 4: 阻塞主线程，直到所有任务完成
	wg.Wait()
	fmt.Println("主线程: 所有协程任务已完成！")
}
func main() {
	task1()
	task2()
}

/*
-------------
预期输出 (任务一的顺序会交替变化):
-------------
--- 任务一：面包生产线 ---
顾客: 准备接收面包...
面包师: 正在制作面包 1...
顾客: 收到面包 1
面包师: 正在制作面包 2...
顾客: 收到面包 2
...
面包师: 下班了 (关闭通道)
--- 任务二：并发打印机 ---
主线程: 等待所有协程结束...
协程 3: 第 1 次打印
协程 1: 第 1 次打印
协程 2: 第 1 次打印
协程 3: 第 2 次打印
... (中间省略) ...
协程 2: 第 3 次打印
主线程: 所有协程任务已完成！
*/
