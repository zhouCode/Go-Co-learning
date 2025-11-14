package main

import (
	"fmt"
)

type Employee struct {
	name, currency string
	salary         float64
}

func main() {
	emp1 := Employee{"Daniel", "$", 20000}
	emp1.PrintSalary()
	printSalary(emp1)
}

// printSalary 方法
func (e Employee) PrintSalary() {
	fmt.Printf("%s 的工资为 %s%.2f\n", e.name, e.currency, e.salary)
}

// printSalary 函数
func printSalary(e Employee) {
	fmt.Printf("%s 的工资为 %s%.2f\n", e.name, e.currency, e.salary)
}
