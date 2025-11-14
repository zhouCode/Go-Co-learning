package main

import "fmt"

//定义结构体 Emp
type Emp struct {
	name string
	age  int
	sex  int
}

func main() {
	//使用 new() 内置函数实例化struct
	emp1 := new(Emp)
	fmt.Printf("emp1:%T ,%v ,%p\n", emp1, emp1, emp1)
	(*emp1).name = "David"
	(*emp1).age = 30
	(*emp1).sex = 1
	//语法糖写法
	emp1.name = "David2"
	emp1.age = 31
	emp1.sex = 0
	fmt.Println(emp1)
	fmt.Printf("emp1:%T ,%v ,%p\n", emp1, emp1, emp1)
}
