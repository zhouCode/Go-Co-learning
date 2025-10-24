package main

// 指针
import "fmt"

func main() {
	// 1. 声明一个普通变量
	name := "Alice"

	// 2. 声明一个指针变量
	// p 的类型是 *string (读作 "string pointer" 或 "指向 string 的指针")
	var p *string

	// 3. 将 name 的地址 赋值给 p
	p = &name

	// 打印 name 的值、p 的值（一个内存地址）
	fmt.Println("Original Name:", name)
	fmt.Printf("Pointer p holds address: %p\n", p)

	// 4. 操作指针：使用 * 来改变它所指向的值
	// *p 意味着 "访问 p 所指向的那个内存地址"
	*p = "Bob" // 这会直接修改变量 name 的值

	fmt.Println("Modified Name:", name)   // 输出 Bob
	fmt.Println("Value via pointer:", *p) // 输出 Bob
}
