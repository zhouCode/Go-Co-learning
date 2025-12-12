package main

import (
	"fmt"
)

func main() {
	s := "我爱Go语言"
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	fmt.Println(len(s))
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x", i, ch)
	}
	fmt.Println()
}
