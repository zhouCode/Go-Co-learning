package main

import "fmt"

func main(){
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println(multiply(2, 3))

	result := func(a, b int) int {
		return a * b
	}(2, 3)
	fmt.Println(result)
	
}