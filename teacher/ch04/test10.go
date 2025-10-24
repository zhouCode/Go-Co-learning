package main

// nil 指针
import "fmt"

func main() {
	var p *int // p 被声明了，但没有指向任何地址，所以 p 是 nil

	fmt.Printf("Pointer p: %v\n", p) // 输出 <nil>

	if p != nil {
		fmt.Println("Value:", *p) // 这行不会执行
	} else {
		fmt.Println("p 是一个 nil 指针")
	}

	// 尝试解引用一个 nil 指针会导致运行时恐慌
	// 下面这行代码会使程序崩溃：
	// panic: runtime error: invalid memory address or nil pointer dereference
	// *p = 10
}
