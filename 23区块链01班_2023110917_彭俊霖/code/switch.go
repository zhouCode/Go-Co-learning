package main

import "fmt"

func main(){
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}