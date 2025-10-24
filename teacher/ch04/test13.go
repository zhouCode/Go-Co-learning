package main

// 指针的指针
import "fmt"

// 这个函数接收一个 "string 指针的指针"
func changePointer(pp **string) {
	// *pp 是 "pp 指向的值"，即 p1 这个指针

	newValue := "New Value"
	// 我们让 p1 指向一个新的地址 (newValue 的地址)
	*pp = &newValue
}

func main() {
	originalValue := "Original Value"

	// p1 指向 originalValue
	p1 := &originalValue

	fmt.Println("p1 指向:", *p1) // 输出 Original Value

	// 我们把 p1 的地址 传给函数
	changePointer(&p1)

	// 现在 p1 本身已经被修改，它指向了一个全新的变量
	fmt.Println("p1 现在指向:", *p1)        // 输出 New Value
	fmt.Println("原始变量:", originalValue) // 输出 Original Value (它没变)
}
