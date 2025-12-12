package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hello", "ll"))
	fmt.Println(strings.Contains("hello", "world"))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Split("a,b,c", ""))
	fmt.Println(strings.Split("a,b,c", "x"))
	fmt.Println(strings.ToLower("Hello"))
	fmt.Println(strings.ToUpper("Hello"))
	fmt.Println(strings.TrimSpace("hello"))
	fmt.Println(strings.EqualFold("hello", "HELLO"))
}
