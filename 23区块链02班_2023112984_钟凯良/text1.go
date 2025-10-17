package main

import "fmt"

func main() {
	score := 85
	if score > 60 {
		fmt.Print("及格了")
	}
	if score < 60 {
		fmt.Print("没有及格")
	}

	if score2 := 55; score2 > 60 {
		fmt.Print("及格了")
	} else {
		fmt.Print("没有及格")
	}
}
