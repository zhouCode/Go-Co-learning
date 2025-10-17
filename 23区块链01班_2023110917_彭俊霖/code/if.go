package main

import "fmt"

func main(){
	score := 85
	if score > 60{
		fmt.Println("及格了")
	} else {
		fmt.Println("不及格")
	}
	if score2 := 85; score2 > 60 {
		fmt.Println("及格了")
	} else {
		fmt.Println("不及格")
	}
}