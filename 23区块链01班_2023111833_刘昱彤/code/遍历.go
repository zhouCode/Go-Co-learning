package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("遍历结束")
	fmt.Println(len(s))
	for i, ch := range s {
		fmt.Println(i, ch)
	}
	fmt.Println("遍历结束")
}
