package main

import (
	"fmt"
)

func main() {
	type Emp struct {
		name string
		age  int
		sex  int
	}
	emp1 := new(Emp)
	fmt.Printf("emp1:%T,%v,%q \n", emp1, emp1, emp1)
	(*emp1).name = "Alice"
	(*emp1).age = 30
	(*emp1).sex = 1
	emp1.name = "Bob"
	emp1.age = 35
	emp1.sex = 0
	fmt.Println(emp1)
}
