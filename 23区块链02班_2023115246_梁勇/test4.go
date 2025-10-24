package main

import "fmt"

//闭包：函数内部定义的函数，并且该函数可以访问其外部函数的变量
func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSequence()
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	newnextInt := intSequence()
	fmt.Println(newnextInt()) //1
	fmt.Println(nextInt())    //？
}

//2025年10月24日14：54
