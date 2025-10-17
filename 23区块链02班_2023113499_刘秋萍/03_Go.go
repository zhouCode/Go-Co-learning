package main

import "fmt"

func main() {
	score := 80
	if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	if score2 := 90; score2 >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
