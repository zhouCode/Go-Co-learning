package main

import (
	"fmt"
)

type Emp struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	emp1 := new(Emp)
	emp1.printSalary()
	printSalary(emp1)
}

func (e *Emp) printSalary() {
	fmt.Printf("%s的工资为10000元\n", e.Name)
}

func printSalary(e *Emp) {
	fmt.Printf("%s的工资为10000元\n", e.Name)
}
