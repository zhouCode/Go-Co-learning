package main

// 指针参数
import "fmt"

func changeName(s string) {
	s = "Charlie" // s 是 main 函数中 name 变量的一个副本
}

func main() {
	name1 := "Alice"
	changeName(name1)
	fmt.Println(name1) // 输出 "Alice"，没有被改变！

	// name2 := "Alice"
	name3 := "Mike"
	ptr := &name3
	// 传递 name 的地址 (&name)
	changeNameByPtr(ptr)
	fmt.Println(name3) // 输出 "Charlie"，成功被改变！
}

// 接收一个 *string (string 指针)
func changeNameByPtr(s *string) {
	// *s 是 "s 指针所指向的值"
	*s = "Charlie"
}
