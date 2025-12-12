package main

import (
	"fmt"
)

func main() {
	type Teacher struct {
		name string
		age  int
		sex  int
	}
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T,%v,%q \n", t1, t1, t1)
	t1.name = "Steven"
	t1.age = 35
	t1.sex = 1
	fmt.Println(t1)
	fmt.Println("-------------------")
}
